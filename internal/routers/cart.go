package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func CartRouter(r *gin.RouterGroup) {
	{
		cartRouter := r.Group("cart").Use(middlewares.Authentication("client"))
		cartHandler := client.NewCartHandler()
		cartRouter.POST("add", cartHandler.Add)
		cartRouter.POST("remove", cartHandler.Remove)
		cartRouter.POST("finalize", cartHandler.Finalize)
	}
}
