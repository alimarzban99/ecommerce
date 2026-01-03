package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/enums"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
)

type CategoryRepository struct {
	*Repository[model.Category, dtoClient.StoreCategoryDTO, dtoClient.UpdateCategoryDTO, client.CategoryResource]
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		&Repository[model.Category, dtoClient.StoreCategoryDTO, dtoClient.UpdateCategoryDTO, client.CategoryResource]{
			database: database.DB(),
		},
	}
}

func (r *CategoryRepository) CategoriesList() ([]client.CategoryPluckResource, error) {
	var categories []client.CategoryPluckResource

	err := r.database.
		Model(&model.Category{}).
		Select("id, title").
		Where("status = ?", enums.Active).
		Scan(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}
