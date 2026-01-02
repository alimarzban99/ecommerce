package model

import (
	"time"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	UserID    uint      `json:"user_id"`
	Type      string    `gorm:"type:varchar(20)" json:"type"` // DEPOSIT, REFUND, PAYMENT
	Amount    float64   `gorm:"not null" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
