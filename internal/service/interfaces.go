package service

import (
	"context"
	authdto "github.com/alimarzban99/ecommerce/internal/dto/auth"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

// AuthServiceInterface defines the interface for authentication service operations
type AuthServiceInterface interface {
	GetVerificationCode(dto *authdto.GetOTPCodeDTO) error
	Verify(dto *authdto.VerifyOTPCodeDTO) (string, error)
	Logout(jti string) error
}

// UserServiceInterface defines the interface for user service operations
type UserServiceInterface interface {
	Profile(id int) (*client.UserResource, error)
	Update(id int, dto *dtoClient.UpdateUserDTO) error
	UpdateProfile(id int, dto *dtoClient.UpdateProfileDTO) (*client.UserResource, error)
}

type UserAddressServiceInterface interface {
	Index(dto dtoClient.ListUserAddressDTO, userId int) (*repository.PaginatedResponse[client.UserAddressListResource], error)
	Show(id, userId int) (*client.UserAddressResource, error)
	Store(dto *dtoClient.StoreUserAddressDTO, userId int) (*client.UserAddressResource, error)
	Update(id, userId int, dto *dtoClient.UpdateUserAddressDTO) error
	Destroy(id, userId int) error
}

// ProductServiceInterface defines the interface for product service operations
type ProductServiceInterface interface {
	List(filter dtoClient.ListProductDTO) (*repository.PaginatedResponse[client.ProductResource], error)
	GetBySlug(slug string) (*client.ProductResource, error)
}

// CategoryServiceInterface defines the interface for category service operations
type CategoryServiceInterface interface {
	CategoriesList() ([]client.CategoryPluckResource, error)
}

// OrderServiceInterface defines the interface for order service operations
type OrderServiceInterface interface {
	List(filter dtoClient.ListOrderDTO) (*repository.PaginatedResponse[client.OrderResource], error)
}

// CartServiceInterface defines the interface for cart service operations
type CartServiceInterface interface {
	Add(ctx context.Context, userID uint, data dtoClient.CartAddDTO) error
	Remove(ctx context.Context, userID uint, data dtoClient.CartRemoveDTO) error
	UpdateQuantity(ctx context.Context, userID uint, data dtoClient.CartUpdateQuantityDTO) error
	Get(ctx context.Context, userID uint) (*client.CartResource, error)
	Finalize(ctx context.Context, userID uint, data dtoClient.CartFinalizeDTO) (*client.OrderResource, error)
}
