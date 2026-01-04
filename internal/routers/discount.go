package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func DiscountRouter(r *gin.RouterGroup, c *container.Container) {
	{
		discountRouter := r.Group("discount").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		discountHandler := client.NewDiscountHandler(c.DiscountService)
		discountRouter.POST("validate", discountHandler.Validate)
	}

}
