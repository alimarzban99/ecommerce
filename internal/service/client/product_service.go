package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type ProductService struct {
	repo repository.ProductRepositoryInterface
}

// NewProductService creates a new product service (kept for backward compatibility)
func NewProductService() *ProductService {
	// This should not be used in production - use NewProductServiceWithDeps instead
	panic("NewProductService() is deprecated. Use NewProductServiceWithDeps() with dependency injection")
}

// NewProductServiceWithDeps creates a new product service with injected dependencies
func NewProductServiceWithDeps(repo repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List(filter dtoClient.ListProductDTO) (*repository.PaginatedResponse[client.ProductResource], error) {
	return s.repo.List(filter)
}

func (s *ProductService) GetBySlug(slug string) (*client.ProductResource, error) {
	return s.repo.FindBySlug(slug)
}
