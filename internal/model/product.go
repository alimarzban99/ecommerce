package model

type Product struct {
	BaseModel
	CategoryID  uint      `gorm:"not null;index" json:"category_id" binding:"required"`
	Name        string    `gorm:"type:varchar(255);not null;index" json:"name" binding:"required,min=3"`
	Slug        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"slug" binding:"required"`
	SKU         string    `gorm:"type:varchar(100);uniqueIndex" json:"sku"`      // Stock Keeping Unit
	Price       int64     `gorm:"not null" json:"price" binding:"required,gt=0"` // Price in smallest currency unit (e.g., cents)
	Stock       int       `gorm:"default:0;not null" json:"stock"`               // Available stock quantity
	Image       string    `gorm:"type:text" json:"image"`
	Description string    `gorm:"type:text" json:"description"`
	Category    *Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:RESTRICT" json:"category,omitempty"`
}

// IsInStock checks if product has available stock
func (p *Product) IsInStock() bool {
	return p.Stock > 0 && p.Status == "active"
}

// CanPurchase checks if product can be purchased
func (p *Product) CanPurchase(quantity int) bool {
	return p.IsInStock() && p.Stock >= quantity
}
