package client

import (
	"github.com/alimarzban99/ecommerce/internal/service/client"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *client.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{service: client.NewCategoryService()}
}

func (h *CategoryHandler) Index(ctx *gin.Context) {

	categories, err := h.service.CategoriesList()

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, categories)
}
