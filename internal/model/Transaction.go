package model

import (
	"github.com/alimarzban99/ecommerce/internal/enums"
)

type Transaction struct {
	BaseModel
	OrderID *uint                 `gorm:"index"` // Nullable, not all transactions are order-related
	Order   *Order                `gorm:"foreignKey:OrderID;constraint:OnDelete:SET NULL"`
	UserID  uint                  `gorm:"not null;index"`
	User    *User                 `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Type    enums.TransactionType `gorm:"type:varchar(20);not null;index"` // deposit, refund, payment
	Amount  float64               `gorm:"not null"`
}
