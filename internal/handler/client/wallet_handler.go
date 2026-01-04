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

type WalletHandler struct {
	service service.WalletServiceInterface
}

func NewWalletHandler(service service.WalletServiceInterface) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) Balance(ctx *gin.Context) {

	userId, err := h.getAuthUser(ctx)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	result, err := h.service.Balance(userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, result)
}

func (h *WalletHandler) Deposit(ctx *gin.Context) {
	dto := new(dtoClient.DepositWalletDTO)
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

	result, err := h.service.Deposit(*dto, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, result)
}

func (h *WalletHandler) getAuthUser(ctx *gin.Context) (int, error) {
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
