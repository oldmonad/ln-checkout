package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
	"github.com/oldmonad/ln-checkout.git/config"
	"github.com/oldmonad/ln-checkout.git/domain"
	"github.com/pkg/errors"
)

type Service struct {
	repository domain.Repository
}

func NewService(repository domain.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreatePaymentLink(data *domain.Payment) (*domain.PaymentResponse, error) {
	return s.repository.CreatePaymentLink(data)
}

func (s *Service) FetchPaymentLink(reference string) (*domain.PaymentResponse, error) {
	paymentLink, _ := s.repository.FetchPaymentLink(reference)

	client, err := config.NewLnConnection()

	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to lnd server")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &lnrpc.Invoice{
		Value: 50,
		Memo:  paymentLink.Description,
	}

	// Send the request to create the invoice
	resp, err := client.AddInvoice(ctx, req)
	if err != nil {
		fmt.Printf("failed to create invoice: %v", err)
		return nil, err
	}

	fmt.Printf("  RHash: %v\n", base64.StdEncoding.EncodeToString(resp.RHash))

	return &domain.PaymentResponse{
		ID:          paymentLink.ID,
		Description: paymentLink.Description,
		Amount:      paymentLink.Amount,
		Reference:   paymentLink.Reference,
		Status:      paymentLink.Status,
		CreatedAt:   paymentLink.CreatedAt,
		Invoice:     resp.PaymentRequest,
		RHash:       base64.StdEncoding.EncodeToString(resp.RHash),
	}, nil
}
