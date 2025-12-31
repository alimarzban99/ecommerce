package model

import (
	"fmt"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	err = db.AutoMigrate(&User{}, &VerificationCode{}, &Token{}, &Post{}, &Category{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("migrate database successfully")
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
