package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
)

type UserRepository struct {
	*Repository[model.User, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.UserResource]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		&Repository[model.User, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.UserResource]{
			database: database.DB(),
		},
	}
}

func (r UserRepository) CheckExistsByMobile(mobile string) (bool, error) {
	var exists bool

	err := r.database.
		Model(&model.User{}).
		Select("count(*) > 0").
		Where("mobile=?", mobile).
		Find(&exists).
		Error

	if err != nil {
		return false, nil
	}

	return exists, nil
}

func (r UserRepository) FindByMobile(mobile string) (*client.UserResource, error) {

	res := &client.UserResource{}
	err := r.database.
		Model(model.User{}).
		Where("mobile=?", mobile).
		Find(res).
		Error

	return res, err
}

func (r UserRepository) UpdateProfile(id int, updateMap map[string]interface{}) error {
	err := r.database.
		Model(&model.User{}).
		Where("id=?", id).
		Updates(updateMap).
		Error

	return err
}
