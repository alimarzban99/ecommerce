package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type TransactionService struct {
	repo repository.TransactionRepositoryInterface
}

func NewTransactionService(repo repository.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) List(dto dtoClient.ListTransactionDTO, userId int) (*repository.PaginatedResponse[client.TransactionListResource], error) {
	return s.repo.List(dto, userId)
}
