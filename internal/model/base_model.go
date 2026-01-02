package model

import (
	"fmt"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Status    string         `gorm:"type:status;default:'active'"`
}

func Starter() {
	db := database.DB()
	err := db.Exec(`
	DO $$ 
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
			CREATE TYPE status AS ENUM ('active', 'inactive', 'blocked');
		END IF;
	END$$;
	`).Error
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(&User{}, &VerificationCode{}, &Token{}, &Product{}, &Category{}, &Order{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("migrate database successfully")
}
