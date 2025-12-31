package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {

	clientRouter := r.Group("client")
	{
		productRouter := clientRouter.Group("product")
		productHandler := client.NewProductHandler()
		productRouter.GET("", productHandler.Index)
		productRouter.PUT(":slug", productHandler.Show)
	}
}
