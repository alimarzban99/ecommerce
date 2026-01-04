package client

import "github.com/alimarzban99/ecommerce/internal/enums"

type TransactionListResource struct {
	ID        uint                  `json:"id"`
	Type      enums.TransactionType `json:"type"`
	Amount    float64               `json:"amount"`
	CreatedAt string                `json:"created_at"`
}
type TransactionResource struct {
	ID        uint                  `json:"id"`
	Type      enums.TransactionType `json:"type"`
	Amount    float64               `json:"amount"`
	CreatedAt string                `json:"created_at"`
}
