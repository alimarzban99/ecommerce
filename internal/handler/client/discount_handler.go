package client

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/internal/service"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	service service.DiscountServiceInterface
}

func NewDiscountHandler(service service.DiscountServiceInterface) *DiscountHandler {
	return &DiscountHandler{service: service}
}

func (h *DiscountHandler) Validate(ctx *gin.Context) {
	dto := new(dtoClient.ValidateDiscountDTO)
	err := ctx.ShouldBindJSON(&dto)

	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	userId, err := h.getAuthUser(ctx)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	result, err := h.service.Validate(dto, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, result)
}

func (h *DiscountHandler) getAuthUser(ctx *gin.Context) (int, error) {
	user, exists := ctx.Get("user")
	if !exists {
		return 0, errors.New("user not found in context")
	}

	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		return 0, errors.New("invalid user type in context")
	}

	return userResource.ID, nil
}
