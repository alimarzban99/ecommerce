package client

import (
	"github.com/alimarzban99/ecommerce/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{repo: repository.NewProductRepository()}
}
