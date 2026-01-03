package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/internal/service"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service service.CartServiceInterface
}

func NewCartHandler(service service.CartServiceInterface) *CartHandler {
	return &CartHandler{service: service}
}

// Add adds a product to the cart
func (h *CartHandler) Add(ctx *gin.Context) {
	dto := new(dtoClient.CartAddDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.ErrorResponse(ctx, "invalid user data")
		return
	}

	err = h.service.Add(ctx.Request.Context(), uint(userResource.ID), *dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, "Product added to cart successfully")
}

func (h *CartHandler) Remove(ctx *gin.Context) {
	dto := new(dtoClient.CartRemoveDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.ErrorResponse(ctx, "invalid user data")
		return
	}

	err = h.service.Remove(ctx.Request.Context(), uint(userResource.ID), *dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, "Product removed from cart successfully")
}

func (h *CartHandler) UpdateQuantity(ctx *gin.Context) {
	dto := new(dtoClient.CartUpdateQuantityDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.ErrorResponse(ctx, "invalid user data")
		return
	}

	err = h.service.UpdateQuantity(ctx.Request.Context(), uint(userResource.ID), *dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, "Cart quantity updated successfully")
}

func (h *CartHandler) Get(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.ErrorResponse(ctx, "invalid user data")
		return
	}

	cart, err := h.service.Get(ctx.Request.Context(), uint(userResource.ID))
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, cart)
}

func (h *CartHandler) Finalize(ctx *gin.Context) {
	dto := new(dtoClient.CartFinalizeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource, ok := user.(*resourceClient.UserResource)

	if !ok {
		response.ErrorResponse(ctx, "invalid user data")
		return
	}

	order, err := h.service.Finalize(ctx.Request.Context(), uint(userResource.ID), *dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, order)
}
