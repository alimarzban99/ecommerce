package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {

	{
		productRouter := r.Group("product")
		productHandler := client.NewProductHandler()
		productRouter.GET("", productHandler.Index)
		productRouter.GET(":slug", productHandler.Show)
	}
}
