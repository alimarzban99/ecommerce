package model

import (
	"github.com/alimarzban99/ecommerce/internal/enums"
	"time"
)

type Order struct {
	ID            uint                `gorm:"primaryKey" json:"id"`
	UserID        uint                `gorm:"not null" json:"user_id"`
	User          *User               `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TrackingCode  string              `gorm:"type:varchar(250)"`
	OrderItems    []OrderItem         `gorm:"foreignKey:OrderID" json:"items"`
	DiscountCode  *string             `json:"discount_code,omitempty"`
	TotalAmount   float64             `gorm:"not null" json:"total_amount"` // مبلغ کل قبل از کسر تخفیف
	FinalAmount   float64             `gorm:"not null" json:"final_amount"` // مبلغ نهایی پس از تخفیف
	Status        enums.OrderStatus   `gorm:"type:varchar(20);default:'pending'" json:"status"`
	PaymentMethod enums.PaymentMethod `gorm:"type:varchar(20)" json:"payment_method"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	CanceledAt    *time.Time          `json:"canceled_at,omitempty"`
	RefundedAt    *time.Time          `json:"refunded_at,omitempty"`
	Transactions  []Transaction       `gorm:"foreignKey:OrderID" json:"transactions,omitempty"` // تراکنش‌های مرتبط با کیف پول
}
