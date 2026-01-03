package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/service"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductServiceInterface
}

func NewProductHandler(service service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) FilterList(c *gin.Context) {}

func (h *ProductHandler) Index(ctx *gin.Context) {
	dto := new(dtoClient.ListProductDTO)
	err := ctx.ShouldBindQuery(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	products, err := h.service.List(*dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, products)
}

func (h *ProductHandler) Show(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		response.ErrorResponse(ctx, "slug is required")
		return
	}

	product, err := h.service.GetBySlug(slug)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, product)
}
