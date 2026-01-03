package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup, c *container.Container) {
	{
		orderRouter := r.Group("order").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		orderHandler := client.NewOrderHandler(c.OrderService)
		orderRouter.GET("", orderHandler.Index)
	}
}
