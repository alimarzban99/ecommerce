package client

import (
	"context"
	"errors"
	"fmt"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/cache"
	"github.com/gin-gonic/gin"
	"time"
)

type CartService struct{}
type CartItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type Cart []CartItem

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) Add(ctx context.Context, userID uint, data dtoClient.CartAddDTO) error {
	key := s.cartKey(userID)

	cart, err := cache.Get[Cart](ctx, key)
	if err != nil {
		cart = Cart{}
	}

	found := false

	for i, item := range cart {
		if item.ProductID == data.ProductId {
			cart[i].Quantity += 1
			found = true
			break
		}
	}

	if !found {
		cart = append(cart, CartItem{
			ProductID: data.ProductId,
			Quantity:  1,
		})
	}

	return cache.Set(ctx, key, cart, 1*time.Hour)
}

func (s *CartService) Remove(ctx context.Context, userID uint, data dtoClient.CartAddDTO) error {

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

			if item.Quantity > 1 {
				item.Quantity--
				newCart = append(newCart, item)
			}
			continue
		}

		newCart = append(newCart, item)
	}

	if !found {
		return errors.New("product not found in cart")
	}

	if len(newCart) == 0 {
		client, _ := cache.Client()
		return client.Del(ctx, key).Err()
	}

	return cache.Set(ctx, key, newCart, 24*time.Hour)
}

func (s *CartService) Finalize(ctx *gin.Context, data dtoClient.CartAddDTO) error {
	user, _ := ctx.Get("user")
	userResource, _ := user.(*resourceClient.UserResource)
	key := s.cartKey(uint(userResource.ID))

	cart, err := cache.Get[map[uint]int](ctx, key)
	if err != nil || len(cart) == 0 {
		return errors.New("cart is empty")
	}

	// TODO:
	// - validate stock
	// - create order
	// - reduce inventory

	client, _ := cache.Client()
	return client.Del(ctx, key).Err()
}

func (s *CartService) cartKey(userID uint) string {
	return fmt.Sprintf("cart:user:%d", userID)
}
