package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func CartRouter(r *gin.RouterGroup, c *container.Container) {
	{
		cartRouter := r.Group("cart").Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		cartHandler := client.NewCartHandler(c.CartService)

		cartRouter.POST("add", cartHandler.Add)
		cartRouter.POST("remove", cartHandler.Remove)
		cartRouter.PUT("update-quantity", cartHandler.UpdateQuantity)
		cartRouter.GET("", cartHandler.Get)
		cartRouter.POST("finalize", cartHandler.Finalize)
	}
}
