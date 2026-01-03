package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index"`
	User      *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Revoked   *bool     `gorm:"default:false" json:"revoked,omitempty"`
	ExpiresAt time.Time `gorm:"null" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

func (token *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if token.ID == "" {
		token.ID = uuid.New().String()
	}
	return
}
