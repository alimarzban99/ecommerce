package client

import (
	"github.com/alimarzban99/ecommerce/internal/service/client"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *client.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{service: client.NewProductService()}
}
func (h *ProductHandler) FilterList(c *gin.Context) {}

func (h *ProductHandler) Index(c *gin.Context) {
}

func (h *ProductHandler) Show(c *gin.Context) {}
