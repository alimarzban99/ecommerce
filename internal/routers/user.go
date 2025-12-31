package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {

	clientRouter := r.Group("user")
	{
		userRouter := clientRouter.Use(middlewares.Authentication("client"))
		userHandler := client.NewUserHandler()
		userRouter.GET("profile", userHandler.Profile)
		userRouter.PUT("", userHandler.Update)
	}
}
