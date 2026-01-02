package model

type OrderItem struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	OrderID   uint     `gorm:"not null" json:"order_id"`
	ProductID uint     `gorm:"not null" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int      `gorm:"not null" json:"quantity"`
	Price     float64  `gorm:"not null" json:"price"` // قیمت هر واحد در زمان سفارش
	Total     float64  `gorm:"not null" json:"total"` // Price * Quantity
}
