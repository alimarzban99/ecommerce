package model

import "time"

type DiscountCode struct {
	BaseModel
	Code            string    `gorm:"unique;not null"`
	UseCount        int       `gorm:"default:0"`
	StartAt         time.Time `gorm:"not null"`
	EndAt           time.Time `gorm:"not null"`
	MinOrderPrice   float64   `gorm:"not null"`
	DiscountPercent float64   `gorm:"not null"`
}
