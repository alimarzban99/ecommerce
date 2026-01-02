package repository

import (
	"errors"
	"github.com/alimarzban99/ecommerce/pkg/converter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
)

type Repository[Model any, CrDTO any, UpDTO any, ResSingle any] struct {
	database *gorm.DB
}

type PaginatedResponse[T any] struct {
	Data            []T   `json:"data"`
	Total           int64 `json:"total"`
	PerPage         int   `json:"per_page"`
	CurrentPage     int   `json:"current_page"`
	LastPage        int   `json:"last_page"`
	From            int   `json:"from"`
	To              int   `json:"to"`
	FirstPage       int   `json:"first_page"`
	HasNextPage     bool  `json:"has_next_page"`
	HasPreviousPage bool  `json:"has_previous_page"`
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) FindOne(id int) (*ResSingle, error) {
	var model Model

	query := r.database.
		Where("id=?", id).
		First(&model)

	if query.Error != nil {
		return nil, query.Error
	}

	if query.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	return converter.TypeConverter[ResSingle](model)
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Create(CreateDTO *CrDTO) (*ResSingle, error) {

	model, _ := converter.TypeConverter[Model](CreateDTO)

	err := r.database.
		Create(&model).
		Error

	if err != nil {
		return nil, err
	}
	return converter.TypeConverter[ResSingle](model)
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Update(id int, UpdateDTO *UpDTO) error {

	updateMap, _ := converter.TypeConverter[map[string]interface{}](UpdateDTO)
	model := new(Model)

	query := r.database.
		Model(model).
		Where("id=?", id).
		Updates(*updateMap)

	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Destroy(id int) error {
	model := new(Model)
	query := r.database.Model(model).Where("id = ?", id).Delete(model)

	if query.Error != nil {
		return query.Error
	}

	if query.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) OrderBY(query *gorm.DB, sort string, direction string) *gorm.DB {

	sortDirection := direction

	query = query.Order(clause.OrderByColumn{
		Column: clause.Column{Name: sort},
		Desc:   sortDirection == "desc",
	})

	return query
}

// Paginate applies pagination to a GORM query and returns paginated results
// Usage: Paginate[ModelType](query, page) or Paginate[ModelType](query, page, limit)
// Default limit is 10, max limit is 100
//
// Example:
//
//	query := db.Model(&Product{}).Where("status = ?", "active")
//	result, err := Paginate[Product](query, 1)        // page 1, 10 per page
//	result, err := Paginate[Product](query, 1, 20)    // page 1, 20 per page
func Paginate[T any](query *gorm.DB, page int, limit ...int) (*PaginatedResponse[T], error) {
	// Default limit is 10
	perPage := 10
	if len(limit) > 0 && limit[0] > 0 {
		perPage = limit[0]
	}

	// Normalize parameters
	if page <= 0 {
		page = 1
	}
	if perPage > 100 {
		perPage = 100
	}

	var items []T
	var total int64

	// Count total records matching the query (before pagination)
	// Use Session to clone the query without affecting the original
	countQuery := query.Session(&gorm.Session{})
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, err
	}

	// Apply pagination to the query
	offset := (page - 1) * perPage
	err := query.Offset(offset).Limit(perPage).Find(&items).Error

	if err != nil {
		return nil, err
	}

	// Calculate pagination metadata
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))
	if lastPage == 0 {
		lastPage = 1
	}

	from := offset + 1
	to := offset + len(items)

	result := &PaginatedResponse[T]{
		Data:            items,
		Total:           total,
		PerPage:         perPage,
		CurrentPage:     page,
		LastPage:        lastPage,
		From:            from,
		To:              to,
		FirstPage:       1,
		HasNextPage:     page < lastPage,
		HasPreviousPage: page > 1,
	}

	return result, nil
}

// Paginate is a convenience method for repositories that works with Model types
func (r *Repository[Model, CrDTO, UpDTO, ResSingle]) Paginate(db *gorm.DB, page int, limit ...int) (*PaginatedResponse[Model], error) {
	return Paginate[Model](db, page, limit...)
}
