package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/service/client"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *client.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: client.NewOrderService()}
}

func (h *OrderHandler) Index(ctx *gin.Context) {
	dto := new(dtoClient.ListOrderDTO)
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
