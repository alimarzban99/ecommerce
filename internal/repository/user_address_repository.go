package repository

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"gorm.io/gorm"
	"time"
)

type UserAddressRepository struct {
	*Repository[model.UserAddress, dtoClient.StoreUserAddressDTO, dtoClient.UpdateUserAddressDTO, client.UserAddressResource]
}

func NewUserAddressRepository() *UserAddressRepository {
	return &UserAddressRepository{
		&Repository[model.UserAddress, dtoClient.StoreUserAddressDTO, dtoClient.UpdateUserAddressDTO, client.UserAddressResource]{
			database: database.DB(),
		},
	}
}

func (r *UserAddressRepository) List(dto dtoClient.ListUserAddressDTO, userId int) (*PaginatedResponse[client.UserAddressListResource], error) {
	limit := dto.Limit
	if limit <= 0 {
		limit = 10
	}

	query := r.database.Model(&model.UserAddress{}).Where("user_id = ?", userId).Order("created_at DESC")

	paginated, err := Paginate[model.UserAddress](query, dto.Page, limit)
	if err != nil {
		return nil, err
	}

	var orderResources []client.UserAddressListResource
	for _, userAdd := range paginated.Data {
		resource := client.UserAddressListResource{
			ID:        userAdd.ID,
			Name:      userAdd.Name,
			Address:   userAdd.Address,
			CreatedAt: userAdd.CreatedAt.Format(time.DateTime),
		}
		orderResources = append(orderResources, resource)
	}

	// Return paginated response with converted resources
	return &PaginatedResponse[client.UserAddressListResource]{
		Data:            orderResources,
		Total:           paginated.Total,
		PerPage:         paginated.PerPage,
		CurrentPage:     paginated.CurrentPage,
		LastPage:        paginated.LastPage,
		From:            paginated.From,
		To:              paginated.To,
		FirstPage:       paginated.FirstPage,
		HasNextPage:     paginated.HasNextPage,
		HasPreviousPage: paginated.HasPreviousPage,
	}, nil
}

func (r *UserAddressRepository) FindOne(id, userId int) (*client.UserAddressResource, error) {
	var address model.UserAddress

	err := r.database.
		Where("id = ? AND user_id = ?", id, userId).
		First(&address).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}

	res := &client.UserAddressResource{
		ID:        address.ID,
		Name:      address.Name,
		City:      address.City,
		Address:   address.Address,
		Status:    address.Status,
		Lat:       address.Lat,
		Lng:       address.Lng,
		CreatedAt: address.CreatedAt.Format(time.DateTime),
	}
	return res, nil
}

func (r *UserAddressRepository) Create(dto *dtoClient.StoreUserAddressDTO, userId int) (*client.UserAddressResource, error) {
	address := model.UserAddress{
		UserID:  userId,
		Name:    dto.Name,
		Address: dto.Address,
		City:    dto.City,
		Status:  dto.Status,
		Lat:     dto.Lat,
		Lng:     dto.Lng,
	}

	err := r.database.Create(&address).Error
	if err != nil {
		return nil, err
	}

	res := &client.UserAddressResource{
		ID:        address.ID,
		Name:      address.Name,
		City:      address.City,
		Address:   address.Address,
		Status:    address.Status,
		Lat:       address.Lat,
		Lng:       address.Lng,
		CreatedAt: address.CreatedAt.Format(time.DateTime),
	}
	return res, nil
}

func (r *UserAddressRepository) Update(id, userId int, dto *dtoClient.UpdateUserAddressDTO) error {
	result := r.database.
		Model(&model.UserAddress{}).
		Where("id = ? AND user_id = ?", id, userId).
		Updates(map[string]any{
			"name":    dto.Name,
			"address": dto.Address,
			"city":    dto.City,
			"status":  dto.Status,
			"lat":     dto.Lat,
			"lng":     dto.Lng,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("address not found")
	}

	return nil
}

func (r *UserAddressRepository) Destroy(id, userId int) error {
	result := r.database.
		Where("id = ? AND user_id = ?", id, userId).
		Delete(&model.UserAddress{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("address not found")
	}

	return nil
}
