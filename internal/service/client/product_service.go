package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{repo: repository.NewProductRepository()}
}

func (s *ProductService) List(filter dtoClient.ListProductDTO) (*repository.PaginatedResponse[client.ProductResource], error) {
	return s.repo.List(filter)
}

func (s *ProductService) GetBySlug(slug string) (*client.ProductResource, error) {
	return s.repo.FindBySlug(slug)
}
