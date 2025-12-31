package client

import (
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{repo: repository.NewCategoryRepository()}
}

func (s *CategoryService) CategoriesList() (*repository.PaginatedResponse[model.Category], error) {
	return s.repo.CategoriesList()
}
