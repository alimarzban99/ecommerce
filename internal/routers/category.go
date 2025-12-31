package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/client"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(r *gin.RouterGroup) {

	clientRouter := r.Group("category")
	{
		categoryHandler := client.NewCategoryHandler()
		clientRouter.GET("", categoryHandler.Index)
	}
}
