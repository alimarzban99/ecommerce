package repository

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
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

func (r *CategoryRepository) CategoriesList() (*PaginatedResponse[model.Category], error) {

	query := r.database.Model(&model.Category{}).
		Select("id, title, created_at").
		Where("status = ?", "active")

	return r.Paginate(query, 1, 50)
}
