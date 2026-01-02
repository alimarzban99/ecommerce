package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	_ "gorm.io/gorm"
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
	// Set default limit if not provided
	limit := filter.Limit
	if limit <= 0 {
		limit = 10
	}

	// Build the query using GORM
	query := r.database.Model(&model.Order{})

	// Sort by newest first (created_at DESC)
	query = query.Order("created_at DESC")

	// Use Laravel-like pagination
	paginated, err := Paginate[model.Order](query, filter.Page, limit)
	if err != nil {
		return nil, err
	}

	// Convert to OrderResource
	var orderResources []client.OrderResource
	for _, order := range paginated.Data {
		resource := client.OrderResource{
			ID:           order.ID,
			TrackingCode: order.TrackingCode,
			Amount:       order.FinalAmount,
			CreatedAt:    order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		orderResources = append(orderResources, resource)
	}

	// Return paginated response with converted resources
	return &PaginatedResponse[client.OrderResource]{
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
