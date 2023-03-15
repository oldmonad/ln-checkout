package repository

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/oldmonad/ln-checkout.git/domain"
	"github.com/oldmonad/ln-checkout.git/queries"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

type Repository struct {
	timeout time.Duration
	queries *queries.Queries
}

func NewRepository(timeout time.Duration, queries *queries.Queries) *Repository {
	return &Repository{
		timeout: timeout,
		queries: queries,
	}
}

func (repo *Repository) CreatePaymentLink(data *domain.Payment) (*domain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	reference := randSeq(10)

	paymentLink, err := repo.queries.CreatePaymentLink(ctx, queries.CreatePaymentLinkParams{
		Description: data.Description,
		Amount:      sql.NullString{String: strconv.Itoa(data.Amount), Valid: true},
		Reference:   sql.NullString{String: reference, Valid: true},
	})

	if err != nil {
		return nil, errors.Wrap(err, "Error Creating a payment link")
	}

	return &domain.PaymentResponse{
		ID:          paymentLink.ID,
		Description: paymentLink.Description,
		Amount:      paymentLink.Amount.String,
		Reference:   paymentLink.Reference.String,
		Status:      paymentLink.Status.String,
		CreatedAt:   paymentLink.CreatedAt.Time.String(),
	}, nil
}

func (repo *Repository) FetchPaymentLink(reference string) (*domain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	defer cancel()

	paymentLink, err := repo.queries.GetPaymentLink(ctx, sql.NullString{String: reference, Valid: true})

	if err != nil {
		return nil, errors.Wrap(err, "Error fetching a payment link")
	}

	return &domain.PaymentResponse{
		ID:          paymentLink.ID,
		Description: paymentLink.Description,
		Amount:      paymentLink.Amount.String,
		Reference:   paymentLink.Reference.String,
		Status:      paymentLink.Status.String,
		CreatedAt:   paymentLink.CreatedAt.Time.String(),
	}, nil
}
