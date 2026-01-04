package client

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/enums"
	"github.com/alimarzban99/ecommerce/internal/repository"
)

type WalletService struct {
	repo repository.TransactionRepositoryInterface
}

func NewWalletService(repo repository.TransactionRepositoryInterface) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Balance(userId int) (float64, error) {
	return s.repo.Balance(userId)
}

func (s *WalletService) Deposit(dto dtoClient.DepositWalletDTO, userId int) (float64, error) {

	if dto.Amount >= 1000000 {
		return 0.0, errors.New("amount too big")
	}

	if dto.Amount <= 100 {
		return 0.0, errors.New("amount too small")
	}

	transaction := dtoClient.StoreTransactionDTO{
		UserID: uint(userId),
		Type:   enums.TransactionDeposit,
		Amount: dto.Amount,
	}
	err := s.repo.Create(&transaction)
	if err != nil {
		return 0.0, err
	}
	return s.repo.Balance(userId)
}
