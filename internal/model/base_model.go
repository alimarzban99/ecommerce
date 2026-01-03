package model

import (
	"github.com/alimarzban99/ecommerce/internal/enums"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Status    enums.Status   `gorm:"type:status;default:'active'" json:"status"`
}
