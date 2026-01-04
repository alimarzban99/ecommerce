package client

import "github.com/alimarzban99/ecommerce/internal/enums"

type ListTransactionDTO struct {
	Page  int `form:"page" binding:"required,min=1"`
	Limit int `form:"limit" binding:"required,min=1"`
}
type StoreTransactionDTO struct {
	UserID uint
	Type   enums.TransactionType
	Amount float64
}
