package client

import (
	"errors"
	"fmt"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type DiscountService struct {
	repo      repository.DiscountRepositoryInterface
	orderRepo repository.OrderRepositoryInterface
}

func NewDiscountService(repo repository.DiscountRepositoryInterface, orderRepo repository.OrderRepositoryInterface) *DiscountService {
	return &DiscountService{
		repo:      repo,
		orderRepo: orderRepo,
	}
}

func (s *DiscountService) Validate(dto *dtoClient.ValidateDiscountDTO, userID int) (*client.DiscountValidationResource, error) {
	discount, err := s.repo.GetCode(dto.Code)
	if err != nil {
		return nil, errors.New("invalid or expired discount code")
	}

	order, err := s.orderRepo.FindOne(int(dto.OrderId))
	fmt.Println(order, discount.MinOrderPrice)
	if order.Amount < discount.MinOrderPrice {
		return nil, errors.New(fmt.Sprintf("Order price must be at least %.2f", discount.MinOrderPrice))
	}

	discountCodeUsed, err := s.orderRepo.CountDiscountCodeUsed(discount.Code, userID)
	if err != nil {
		return nil, err
	}
	if discount.UseCount >= *discountCodeUsed {
		return nil, errors.New("discount code usage limit reached")
	}

	discountAmount := (order.Amount * discount.DiscountPercent) / 100

	return &client.DiscountValidationResource{
		FinalPrice:     order.Amount - discountAmount,
		DiscountAmount: discountAmount,
	}, nil

}
