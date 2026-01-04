package repository

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/enums"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/database"
	_ "gorm.io/gorm"
	"time"
)

type TransactionRepository struct {
	*Repository[model.Transaction, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.TransactionResource]
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		&Repository[model.Transaction, dtoClient.StoreUserDTO, dtoClient.UpdateUserDTO, client.TransactionResource]{
			database: database.DB(),
		},
	}
}

func (r *TransactionRepository) List(dto dtoClient.ListTransactionDTO, userId int) (*PaginatedResponse[client.TransactionListResource], error) {
	limit := dto.Limit
	if limit <= 0 {
		limit = 10
	}

	query := r.database.Model(&model.Transaction{}).Where("user_id = ?", userId).Order("created_at DESC")

	paginated, err := Paginate[model.Transaction](query, dto.Page, limit)
	if err != nil {
		return nil, err
	}

	var transactionResources []client.TransactionListResource
	for _, transaction := range paginated.Data {
		resource := client.TransactionListResource{
			ID:        transaction.ID,
			Type:      transaction.Type,
			Amount:    transaction.Amount,
			CreatedAt: transaction.CreatedAt.Format(time.DateTime),
		}
		transactionResources = append(transactionResources, resource)
	}

	return &PaginatedResponse[client.TransactionListResource]{
		Data:            transactionResources,
		Total:           paginated.Total,
		PerPage:         paginated.PerPage,
		CurrentPage:     paginated.CurrentPage,
		LastPage:        paginated.LastPage,
		From:            paginated.From,
		To:              paginated.To,
		FirstPage:       paginated.FirstPage,
		HasNextPage:     paginated.HasNextPage,
		HasPreviousPage: paginated.HasPreviousPage,
	}, nil
}

func (r *TransactionRepository) Balance(userId int) (float64, error) {
	var balance float64

	err := r.database.Raw(`
        SELECT 
            (SELECT COALESCE(SUM(amount), 0) 
             FROM transactions 
             WHERE user_id = ? 
             AND type IN (?, ?) 
             AND deleted_at IS NULL) 
             -
            (SELECT COALESCE(SUM(amount), 0) 
             FROM transactions 
             WHERE user_id = ? 
             AND type = ? 
             AND deleted_at IS NULL) AS balance`,
		userId,
		enums.TransactionDeposit,
		enums.TransactionRefund,
		userId,
		enums.TransactionPayment).Scan(&balance).Error

	if err != nil {
		return 0, errors.New("failed to get balance")
	}

	return balance, nil
}

func (r *TransactionRepository) Create(dto *dtoClient.StoreTransactionDTO) error {
	transaction := model.Transaction{
		UserID: dto.UserID,
		Type:   dto.Type,
		Amount: dto.Amount,
	}

	err := r.database.Create(&transaction).Error
	if err != nil {
		return err
	}

	return nil
}
