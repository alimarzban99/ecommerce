package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	_ "gorm.io/gorm"
	"strings"
)

type ProductRepository struct {
	*Repository[model.Product, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.ProductResource]
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		&Repository[model.Product, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.ProductResource]{
			database: database.DB(),
		},
	}
}

func (r *ProductRepository) List(filter dtoClient.ListProductDTO) (*PaginatedResponse[client.ProductResource], error) {

	query := r.database.Model(&model.Product{})

	// Preload category relationship
	query = query.Preload("Category")

	// Apply filters
	if filter.Search != "" {
		searchTerm := "%" + strings.ToLower(strings.TrimSpace(filter.Search)) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(slug) LIKE ?", searchTerm, searchTerm)
	}

	if filter.Category > 0 {
		query = query.Where("category_id = ?", filter.Category)
	}

	// Filter only active products
	query = query.Where("status = ?", "active")

	// Sort by newest first (created_at DESC)
	query = query.Order("created_at DESC")

	// Use Laravel-like pagination
	paginated, err := Paginate[model.Product](query, filter.Page, filter.Limit)
	if err != nil {
		return nil, err
	}

	// Convert to ProductResource
	var productResources []client.ProductResource
	for _, product := range paginated.Data {
		resource := client.ProductResource{
			ID:          product.ID,
			Name:        product.Name,
			Slug:        product.Slug,
			Image:       product.Image,
			Description: product.Description,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if product.Category != nil {
			resource.Category = &client.CategoryResource{
				ID:    product.Category.ID,
				Title: product.Category.Title,
			}
		}

		productResources = append(productResources, resource)
	}

	// Return paginated response with converted resources
	return &PaginatedResponse[client.ProductResource]{
		Data:            productResources,
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

func (r *ProductRepository) FindBySlug(slug string) (*client.ProductResource, error) {
	var product model.Product

	err := r.database.
		Preload("Category").
		Where("slug = ? AND status = ?", slug, "active").
		First(&product).Error

	if err != nil {
		return nil, err
	}

	res := &client.ProductResource{
		ID:          product.ID,
		Name:        product.Name,
		Slug:        product.Slug,
		Image:       product.Image,
		Description: product.Description,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt.Format("2006-01-01T15:04:05"),
	}

	if product.CategoryID != 0 {
		res.Category = &client.CategoryResource{
			ID:    product.Category.ID,
			Title: product.Category.Title,
		}
	}

	return res, nil
}
