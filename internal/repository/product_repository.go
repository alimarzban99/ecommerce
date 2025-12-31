package repository

import (
	dtoAdmin "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
)

type ProductRepository struct {
	*Repository[model.User, dtoAdmin.StoreUserDTO, dtoAdmin.UpdateUserDTO, client.UserResource]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		&Repository[model.User, dtoAdmin.StoreUserDTO, dtoAdmin.UpdateUserDTO, client.UserResource]{
			database: database.DB(),
		},
	}
}
