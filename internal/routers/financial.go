package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup, c *container.Container) {
	{
		orderRouter := r.Group("financial/order").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		orderHandler := client.NewOrderHandler(c.OrderService)
		orderRouter.GET("", orderHandler.Index)
	}

	{
		transactionRouter := r.Group("financial/transaction").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		transactionHandler := client.NewTransactionHandler(c.TransactionService)
		transactionRouter.GET("", transactionHandler.Index)
	}

	{
		walletRouter := r.Group("financial/wallet").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		walletHandler := client.NewWalletHandler(c.WalletService)
		walletRouter.GET("", walletHandler.Balance)
		walletRouter.POST("deposit", walletHandler.Deposit)
	}
}
