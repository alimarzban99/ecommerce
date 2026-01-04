package repository

import (
	authdto "github.com/alimarzban99/ecommerce/internal/dto/auth"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/model"
	authResources "github.com/alimarzban99/ecommerce/internal/resources/auth"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

// UserRepositoryInterface defines the interface for user repository operations
type UserRepositoryInterface interface {
	FindOne(id int) (*client.UserResource, error)
	Create(dto *dtoClient.StoreUserDTO) (*client.UserResource, error)
	Update(id int, dto *dtoClient.UpdateUserDTO) error
	Destroy(id int) error
	CheckExistsByMobile(mobile string) (bool, error)
	FindByMobile(mobile string) (*client.UserResource, error)
	UpdateProfile(id int, updateMap map[string]interface{}) error
}

type UserAddressRepositoryInterface interface {
	List(dto dtoClient.ListUserAddressDTO, userId int) (*PaginatedResponse[client.UserAddressListResource], error)
	FindOne(id, userId int) (*client.UserAddressResource, error)
	Create(dto *dtoClient.StoreUserAddressDTO, userId int) (*client.UserAddressResource, error)
	Update(id, userId int, dto *dtoClient.UpdateUserAddressDTO) error
	Destroy(id, userId int) error
}

// TokenRepositoryInterface defines the interface for token repository operations
type TokenRepositoryInterface interface {
	Create(dto *authdto.TokenCreate) (*authResources.TokenResponse, error)
	FindToken(jti string) (bool, error)
	ExpiredToken(jti string) error
}

// VerificationCodeRepositoryInterface defines the interface for verification code repository operations
type VerificationCodeRepositoryInterface interface {
	Create(dto *authdto.CreateOTPCodeDTO) (*authResources.CodeResponse, error)
	ValidCode(dto *authdto.VerifyOTPCodeDTO) (bool, error)
}

// ProductRepositoryInterface defines the interface for product repository operations
type ProductRepositoryInterface interface {
	FindOne(id int) (*client.ProductResource, error)
	Create(dto *dtoClient.StoreUserDTO) (*client.ProductResource, error)
	Update(id int, dto *dtoClient.UpdateUserDTO) error
	Destroy(id int) error
	List(filter dtoClient.ListProductDTO) (*PaginatedResponse[client.ProductResource], error)
	FindBySlug(slug string) (*client.ProductResource, error)
	FindByID(id uint) (*model.Product, error)
}

// CategoryRepositoryInterface defines the interface for category repository operations
type CategoryRepositoryInterface interface {
	FindOne(id int) (*client.CategoryResource, error)
	Create(dto *dtoClient.StoreCategoryDTO) (*client.CategoryResource, error)
	Update(id int, dto *dtoClient.UpdateCategoryDTO) error
	Destroy(id int) error
	CategoriesList() ([]client.CategoryPluckResource, error)
}

type OrderRepositoryInterface interface {
	FindOne(id int) (*client.OrderResource, error)
	Create(dto *dtoClient.StoreUserDTO) (*client.OrderResource, error)
	Update(id int, dto *dtoClient.UpdateUserDTO) error
	Destroy(id int) error
	List(filter dtoClient.ListOrderDTO) (*PaginatedResponse[client.OrderResource], error)
	CountDiscountCodeUsed(code string, userId int) (*int, error)
}

type TransactionRepositoryInterface interface {
	List(dto dtoClient.ListTransactionDTO, userId int) (*PaginatedResponse[client.TransactionListResource], error)
	Balance(userId int) (float64, error)
	Create(dto *dtoClient.StoreTransactionDTO) error
}

type DiscountRepositoryInterface interface {
	GetCode(code string) (*model.DiscountCode, error)
}
