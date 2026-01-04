package container

import (
	"crypto/rsa"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/service"
	serviceAuth "github.com/alimarzban99/ecommerce/internal/service/auth"
	serviceClient "github.com/alimarzban99/ecommerce/internal/service/client"
	"gorm.io/gorm"
)

// Container holds all dependencies for the application
type Container struct {
	// Database
	DB *gorm.DB

	// JWT Keys
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	// Repositories
	UserRepository             repository.UserRepositoryInterface
	UserAddressRepository      repository.UserAddressRepositoryInterface
	TokenRepository            repository.TokenRepositoryInterface
	VerificationCodeRepository repository.VerificationCodeRepositoryInterface
	ProductRepository          repository.ProductRepositoryInterface
	CategoryRepository         repository.CategoryRepositoryInterface
	OrderRepository            repository.OrderRepositoryInterface

	// Services
	AuthService        service.AuthServiceInterface
	UserService        service.UserServiceInterface
	UserAddressService service.UserAddressServiceInterface
	ProductService     service.ProductServiceInterface
	CategoryService    service.CategoryServiceInterface
	OrderService       service.OrderServiceInterface
	CartService        service.CartServiceInterface
}

// NewContainer creates a new dependency injection container
func NewContainer(db *gorm.DB, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *Container {
	// Initialize repositories
	userRepo := repository.NewUserRepository()
	userAddressRepo := repository.NewUserAddressRepository()
	tokenRepo := repository.NewTokenRepository()
	verificationCodeRepo := repository.NewVerificationCodeRepository()
	productRepo := repository.NewProductRepository()
	categoryRepo := repository.NewCategoryRepository()
	orderRepo := repository.NewOrderRepository()

	// Initialize services with injected repositories
	authService := serviceAuth.NewAuthServiceWithDeps(
		verificationCodeRepo,
		tokenRepo,
		userRepo,
		privateKey,
	)
	userService := serviceClient.NewUserServiceWithDeps(userRepo)
	userAddressService := serviceClient.NewUserAddressService(userAddressRepo)
	productService := serviceClient.NewProductServiceWithDeps(productRepo)
	categoryService := serviceClient.NewCategoryServiceWithDeps(categoryRepo)
	orderService := serviceClient.NewOrderServiceWithDeps(orderRepo)
	cartService := serviceClient.NewCartServiceWithDeps(productRepo, orderRepo)

	return &Container{
		DB:                         db,
		PrivateKey:                 privateKey,
		PublicKey:                  publicKey,
		UserRepository:             userRepo,
		TokenRepository:            tokenRepo,
		VerificationCodeRepository: verificationCodeRepo,
		ProductRepository:          productRepo,
		CategoryRepository:         categoryRepo,
		OrderRepository:            orderRepo,
		AuthService:                authService,
		UserService:                userService,
		UserAddressService:         userAddressService,
		ProductService:             productService,
		CategoryService:            categoryService,
		OrderService:               orderService,
		CartService:                cartService,
	}
}
