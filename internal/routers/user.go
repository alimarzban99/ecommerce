package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, c *container.Container) {

	clientRouter := r.Group("user")
	{
		userRouter := clientRouter.Use(middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"client",
		))
		userHandler := client.NewUserHandler(c.UserService)
		userRouter.GET("profile", userHandler.Profile)
		userRouter.PUT("", userHandler.Update)

	}

	{
		userAddressRouter := clientRouter.Group("address").
			Use(middlewares.Authentication(
				c.PublicKey,
				c.TokenRepository,
				c.UserRepository,
				"client",
			))
		userAddressHandler := client.NewUserAddressHandler(c.UserAddressService)
		userAddressRouter.GET("", userAddressHandler.Index)
		userAddressRouter.GET(":id", userAddressHandler.Show)
		userAddressRouter.POST("", userAddressHandler.Store)
		userAddressRouter.PUT(":id", userAddressHandler.Update)
		userAddressRouter.DELETE(":id", userAddressHandler.Destroy)
	}
}
