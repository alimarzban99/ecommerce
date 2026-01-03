package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(r *gin.RouterGroup, c *container.Container) {

	clientRouter := r.Group("category")
	{
		categoryHandler := client.NewCategoryHandler(c.CategoryService)
		clientRouter.GET("", categoryHandler.Index)
	}
}
