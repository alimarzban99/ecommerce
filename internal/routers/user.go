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
}
