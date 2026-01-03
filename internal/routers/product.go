package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup, c *container.Container) {

	{
		productRouter := r.Group("product")
		productHandler := client.NewProductHandler(c.ProductService)
		productRouter.GET("", productHandler.Index)
		productRouter.GET(":slug", productHandler.Show)
	}
}
