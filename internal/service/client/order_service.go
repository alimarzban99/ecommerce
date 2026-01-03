package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type OrderService struct {
	repo repository.OrderRepositoryInterface
}

// NewOrderService creates a new order service (kept for backward compatibility)
func NewOrderService() *OrderService {
	// This should not be used in production - use NewOrderServiceWithDeps instead
	panic("NewOrderService() is deprecated. Use NewOrderServiceWithDeps() with dependency injection")
}

// NewOrderServiceWithDeps creates a new order service with injected dependencies
func NewOrderServiceWithDeps(repo repository.OrderRepositoryInterface) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) List(filter dtoClient.ListOrderDTO) (*repository.PaginatedResponse[client.OrderResource], error) {
	return s.repo.List(filter)
}
