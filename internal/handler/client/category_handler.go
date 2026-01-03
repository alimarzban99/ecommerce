package client

import (
	"github.com/alimarzban99/ecommerce/internal/service"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service service.CategoryServiceInterface
}

func NewCategoryHandler(service service.CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) Index(ctx *gin.Context) {

	categories, err := h.service.CategoriesList()

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, categories)
}
