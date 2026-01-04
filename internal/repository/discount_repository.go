package repository

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/enums"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type DiscountRepository struct {
	*Repository[model.DiscountCode, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.OrderResource]
}

func NewDiscountRepository() *DiscountRepository {
	return &DiscountRepository{
		&Repository[model.DiscountCode, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.OrderResource]{
			database: database.DB(),
		},
	}
}

func (r *DiscountRepository) GetCode(code string) (*model.DiscountCode, error) {

	var discount model.DiscountCode

	err := r.database.
		Where("code = ? AND status = ? AND start_at <= ? AND end_at >= ?", code, enums.Active, time.Now(), time.Now()).
		First(&discount).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}

	return &discount, nil
}
