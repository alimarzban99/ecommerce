package model

type OrderItem struct {
	BaseModel
	OrderID   uint     `gorm:"not null;index"`
	Order     *Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	ProductID uint     `gorm:"not null;index"`
	Product   *Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
	Quantity  int      `gorm:"not null"`
	Price     float64  `gorm:"not null"`
	Total     float64  `gorm:"not null"`
}
