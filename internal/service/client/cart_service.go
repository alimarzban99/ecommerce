package client

import (
	"context"
	"errors"
	"fmt"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/enums"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/repository"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/cache"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CartService struct {
	productRepo *repository.ProductRepository
	orderRepo   *repository.OrderRepository
}

type CartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type Cart []CartItem

func NewCartService() *CartService {
	return &CartService{
		productRepo: repository.NewProductRepository(),
		orderRepo:   repository.NewOrderRepository(),
	}
}

func (s *CartService) Add(ctx context.Context, userID uint, data dtoClient.CartAddDTO) error {
	// Validate product exists and is in stock
	product, err := s.productRepo.FindByID(data.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if !product.IsInStock() {
		return errors.New("product is out of stock")
	}

	key := s.cartKey(userID)
	cart, err := cache.Get[Cart](ctx, key)
	if err != nil {
		cart = Cart{}
	}

	found := false
	for i, item := range cart {
		if item.ProductID == data.ProductId {
			// Check if adding one more would exceed stock
			if !product.CanPurchase(item.Quantity + 1) {
				return fmt.Errorf("cannot add more items. Only %d available in stock", product.Stock)
			}
			cart[i].Quantity += 1
			found = true
			break
		}
	}

	if !found {
		if !product.CanPurchase(1) {
			return fmt.Errorf("product is out of stock. Only %d available", product.Stock)
		}
		cart = append(cart, CartItem{
			ProductID: data.ProductId,
			Quantity:  1,
		})
	}

	return cache.Set(ctx, key, cart, 24*time.Hour)
}

func (s *CartService) Remove(ctx context.Context, userID uint, data dtoClient.CartRemoveDTO) error {
	key := s.cartKey(userID)

	cart, err := cache.Get[Cart](ctx, key)
	if err != nil || len(cart) == 0 {
		return errors.New("cart is empty")
	}

	found := false
	newCart := make(Cart, 0, len(cart))

	for _, item := range cart {
		if item.ProductID == data.ProductId {
			found = true
			continue // Skip this item
		}
		newCart = append(newCart, item)
	}

	if !found {
		return errors.New("product not found in cart")
	}

	if len(newCart) == 0 {
		redisClient, _ := cache.Client()
		return redisClient.Del(ctx, key).Err()
	}

	return cache.Set(ctx, key, newCart, 24*time.Hour)
}

func (s *CartService) UpdateQuantity(ctx context.Context, userID uint, data dtoClient.CartUpdateQuantityDTO) error {
	// Validate product exists and stock is available
	product, err := s.productRepo.FindByID(data.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	if !product.CanPurchase(data.Quantity) {
		return fmt.Errorf("insufficient stock. Only %d available", product.Stock)
	}

	key := s.cartKey(userID)
	cart, err := cache.Get[Cart](ctx, key)
	if err != nil || len(cart) == 0 {
		return errors.New("cart is empty")
	}

	found := false
	for i, item := range cart {
		if item.ProductID == data.ProductId {
			cart[i].Quantity = data.Quantity
			found = true
			break
		}
	}

	if !found {
		return errors.New("product not found in cart")
	}

	return cache.Set(ctx, key, cart, 24*time.Hour)
}

func (s *CartService) Get(ctx context.Context, userID uint) (*resourceClient.CartResource, error) {
	key := s.cartKey(userID)

	cart, err := cache.Get[Cart](ctx, key)
	if err != nil || len(cart) == 0 {
		return &resourceClient.CartResource{
			Items:      []resourceClient.CartItemResource{},
			TotalItems: 0,
			Subtotal:   0,
		}, nil
	}

	var items []resourceClient.CartItemResource
	var totalItems int
	var subtotal int64

	for _, item := range cart {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			continue
		}

		itemSubtotal := product.Price * int64(item.Quantity)

		cartItem := resourceClient.CartItemResource{
			ProductID:   product.ID,
			ProductName: product.Name,
			ProductSlug: product.Slug,
			Image:       product.Image,
			Price:       product.Price,
			Quantity:    item.Quantity,
			Subtotal:    itemSubtotal,
		}

		items = append(items, cartItem)
		totalItems += item.Quantity
		subtotal += itemSubtotal
	}

	return &resourceClient.CartResource{
		Items:      items,
		TotalItems: totalItems,
		Subtotal:   subtotal,
	}, nil
}

func (s *CartService) Finalize(ctx context.Context, userID uint, data dtoClient.CartFinalizeDTO) (*resourceClient.OrderResource, error) {
	key := s.cartKey(userID)

	cart, err := cache.Get[Cart](ctx, key)
	if err != nil || len(cart) == 0 {
		return nil, errors.New("cart is empty")
	}

	// Validate payment method
	paymentMethod := enums.PaymentMethod(data.PaymentMethod)
	if paymentMethod != enums.PaymentGateway && paymentMethod != enums.PaymentWallet {
		return nil, errors.New("invalid payment method")
	}

	// Start database transaction
	db := database.DB()
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Build order items and validate stock
	var orderItems []model.OrderItem
	var totalAmount float64

	for _, cartItem := range cart {
		product, err := s.productRepo.FindByID(cartItem.ProductID)
		if err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("product with ID %d not found", cartItem.ProductID)
			}
			return nil, err
		}

		// Validate stock
		if !product.CanPurchase(cartItem.Quantity) {
			tx.Rollback()
			return nil, fmt.Errorf("insufficient stock for product %s. Only %d available", product.Name, product.Stock)
		}

		// Calculate prices (convert from int64 cents to float64)
		price := float64(product.Price)
		itemTotal := price * float64(cartItem.Quantity)
		totalAmount += itemTotal

		orderItems = append(orderItems, model.OrderItem{
			ProductID: product.ID,
			Quantity:  cartItem.Quantity,
			Price:     price,
			Total:     itemTotal,
		})

		// Reduce stock
		product.Stock -= cartItem.Quantity
		if err := tx.Model(&model.Product{}).Where("id = ?", product.ID).Update("stock", product.Stock).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Generate tracking code
	trackingCode := uuid.New().String()

	// Create order
	order := model.Order{
		UserID:        userID,
		TrackingCode:  trackingCode,
		OrderItems:    orderItems,
		DiscountCode:  data.DiscountCode,
		TotalAmount:   totalAmount,
		FinalAmount:   totalAmount, // TODO: Apply discount if discount code is provided
		Status:        enums.OrderPending,
		PaymentMethod: paymentMethod,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Clear cart from Redis
	redisClient, _ := cache.Client()
	if redisClient != nil {
		redisClient.Del(ctx, key)
	}

	return &resourceClient.OrderResource{
		ID:           order.ID,
		TrackingCode: order.TrackingCode,
		Amount:       order.FinalAmount,
		CreatedAt:    order.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *CartService) cartKey(userID uint) string {
	return fmt.Sprintf("cart:user:%d", userID)
}
