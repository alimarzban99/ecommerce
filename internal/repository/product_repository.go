package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	_ "gorm.io/gorm"
	"math"
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
	query := r.database.Model(&model.Product{})

	// Preload category relationship
	query = query.Preload("Category")

	//Apply filters
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

	// Get total count before pagination
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination
	offset := (filter.Page - 1) * filter.Limit
	query = query.Offset(offset).Limit(filter.Limit)

	// Execute query and get results
	var posts []model.Product
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	// Convert to ProductResource
	var products []client.ProductResource
	for _, post := range posts {
		product := client.ProductResource{
			ID:          post.ID,
			Name:        post.Name,
			Slug:        post.Slug,
			Image:       post.Image,
			Description: post.Description,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt.Format("2006-01-01T15:04:05"),
		}

		if post.Description != "" {
			product.Description = post.Description
		}

		if post.CategoryID != 0 {
			product.Category = &client.CategoryResource{
				ID:    post.Category.ID,
				Title: post.Category.Title,
			}
		}

		products = append(products, product)
	}

	// Calculate pagination metadata
	lastPage := int(math.Ceil(float64(total) / float64(filter.Limit)))
	if lastPage == 0 {
		lastPage = 1
	}

	// Build paginated response
	result := &PaginatedResponse[client.ProductResource]{
		Data:            products,
		Total:           total,
		PerPage:         filter.Limit,
		CurrentPage:     filter.Page,
		LastPage:        lastPage,
		From:            offset + 1,
		To:              offset + len(products),
		FirstPage:       1,
		HasNextPage:     filter.Page < lastPage,
		HasPreviousPage: filter.Page > 1,
	}

	return result, nil
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
