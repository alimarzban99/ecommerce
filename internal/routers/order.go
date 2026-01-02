package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	{
		orderRouter := r.Group("order")
		orderHandler := client.NewOrderHandler()
		orderRouter.GET("", orderHandler.Index)
	}
}
