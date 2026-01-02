package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	_ "gorm.io/gorm"
	"math"
)

type OrderRepository struct {
	*Repository[model.Order, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.OrderResource]
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		&Repository[model.Order, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.OrderResource]{
			database: database.DB(),
		},
	}
}

func (r *OrderRepository) List(filter dtoClient.ListOrderDTO) (*PaginatedResponse[client.OrderResource], error) {
	// Set default values
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}

	// Build the query using GORM
	query := r.database.Model(&model.Order{})

	query = query.Order("created_at DESC")

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.Limit
	query = query.Offset(offset).Limit(filter.Limit)

	// Execute query and get results
	var posts []model.Order
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	// Convert to OrderResource
	var orders []client.OrderResource
	for _, post := range posts {
		order := client.OrderResource{
			ID:           post.ID,
			TrackingCode: post.TrackingCode,
			Amount:       post.FinalAmount,
			CreatedAt:    post.CreatedAt.Format("2006-01-01T15:04:05"),
		}

		orders = append(orders, order)
	}

	// Calculate pagination metadata
	lastPage := int(math.Ceil(float64(total) / float64(filter.Limit)))
	if lastPage == 0 {
		lastPage = 1
	}

	// Build paginated response
	result := &PaginatedResponse[client.OrderResource]{
		Data:            orders,
		Total:           total,
		PerPage:         filter.Limit,
		CurrentPage:     filter.Page,
		LastPage:        lastPage,
		From:            offset + 1,
		To:              offset + len(orders),
		FirstPage:       1,
		HasNextPage:     filter.Page < lastPage,
		HasPreviousPage: filter.Page > 1,
	}

	return result, nil
}
