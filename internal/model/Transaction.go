package model

import (
	"github.com/alimarzban99/ecommerce/internal/enums"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID        uint                  `gorm:"primarykey" json:"id"`
	OrderID   *uint                 `gorm:"index"` // Nullable, not all transactions are order-related
	Order     *Order                `gorm:"foreignKey:OrderID;constraint:OnDelete:SET NULL"`
	UserID    uint                  `gorm:"not null;index"`
	User      *User                 `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Type      enums.TransactionType `gorm:"type:varchar(20);not null;index"` // deposit, refund, payment
	Amount    float64               `gorm:"not null"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt gorm.DeletedAt        `gorm:"index" json:"-"`
}
