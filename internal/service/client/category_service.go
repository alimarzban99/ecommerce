package client

import (
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type CategoryService struct {
	repo repository.CategoryRepositoryInterface
}

// NewCategoryService creates a new category service (kept for backward compatibility)
func NewCategoryService() *CategoryService {
	// This should not be used in production - use NewCategoryServiceWithDeps instead
	panic("NewCategoryService() is deprecated. Use NewCategoryServiceWithDeps() with dependency injection")
}

// NewCategoryServiceWithDeps creates a new category service with injected dependencies
func NewCategoryServiceWithDeps(repo repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CategoriesList() ([]client.CategoryPluckResource, error) {
	return s.repo.CategoriesList()
}
