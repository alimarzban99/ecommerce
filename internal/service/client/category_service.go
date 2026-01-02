package client

import (
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepository()}
}

func (s *CategoryService) CategoriesList() ([]client.CategoryPluckResource, error) {
	return s.repo.CategoriesList()
}
