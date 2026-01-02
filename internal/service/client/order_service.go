package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{repo: repository.NewOrderRepository()}
}

func (s *OrderService) List(filter dtoClient.ListOrderDTO) (*repository.PaginatedResponse[client.OrderResource], error) {
	return s.repo.List(filter)
}
