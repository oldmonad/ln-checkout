package domain

type Service interface {
	CreatePaymentLink(data *Payment) (*PaymentResponse, error)
	FetchPaymentLink(reference string) (*PaymentResponse, error)
}

type Repository interface {
	CreatePaymentLink(data *Payment) (*PaymentResponse, error)
	FetchPaymentLink(reference string) (*PaymentResponse, error)
}
