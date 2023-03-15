package domain

type Payment struct {
	Description string `json:"description" bson:"description"`
	Amount      int    `json:"amount"`
}

type PaymentResponse struct {
	ID          int64
	Description string `json:"description" bson:"description"`
	Amount      string `json:"amount"`
	Reference   string `json:"reference" bson:"reference"`
	Status      string `json:"status" bson:"status"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	Invoice     string `json:"invoice" bson:"invoice"`
	RHash       string `json:"rhash" bson:"rhash"`
}

type PaymentWithInvoiceResponse struct {
	ID          int64
	Description string `json:"description" bson:"description"`
	Amount      string `json:"amount"`
	Reference   string `json:"reference" bson:"reference"`
	Status      string `json:"status" bson:"status"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	Invoice     string `json:"invoice" bson:"invoice"`
	RHash       string `json:"rhash" bson:"rhash"`
}
