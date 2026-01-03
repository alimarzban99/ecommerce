package model

import (
	"github.com/alimarzban99/ecommerce/internal/enums"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID            uint                `gorm:"primaryKey"`
	UserID        uint                `gorm:"not null;index"`
	User          *User               `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	TrackingCode  string              `gorm:"type:varchar(250)"`
	OrderItems    []OrderItem         `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	DiscountCode  *string             `gorm:"type:varchar(50)"`
	TotalAmount   float64             `gorm:"not null"`
	FinalAmount   float64             `gorm:"not null"`
	Status        enums.OrderStatus   `gorm:"type:order_status;default:'pending';index"`
	PaymentMethod enums.PaymentMethod `gorm:"type:payment_method;index"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	DeletedAt     gorm.DeletedAt      `gorm:"index" json:"-"`
	CanceledAt    *time.Time
	RefundedAt    *time.Time
	Transactions  []Transaction `gorm:"foreignKey:OrderID;constraint:OnDelete:SET NULL"`
}
