package client

import (
	"errors"
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/internal/service"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserAddressHandler struct {
	service service.UserAddressServiceInterface
}

func NewUserAddressHandler(service service.UserAddressServiceInterface) *UserAddressHandler {
	return &UserAddressHandler{service: service}
}

func (h *UserAddressHandler) Index(ctx *gin.Context) {
	dto := new(dtoClient.ListUserAddressDTO)
	err := ctx.ShouldBindQuery(&dto)

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

	result, err := h.service.Index(*dto, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, result)
}
func (h *UserAddressHandler) Show(ctx *gin.Context) {
	userId, err := h.getAuthUser(ctx)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	idParam := ctx.Param("id")
	userAddressId, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, "invalid user address")
		return
	}

	result, err := h.service.Show(userAddressId, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.SuccessResponse(ctx, result)
}
func (h *UserAddressHandler) Store(ctx *gin.Context) {
	dto := new(dtoClient.StoreUserAddressDTO)
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

	result, err := h.service.Store(dto, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, result)
}
func (h *UserAddressHandler) Update(ctx *gin.Context) {
	dto := new(dtoClient.UpdateUserAddressDTO)
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

	idParam := ctx.Param("id")
	userAddressId, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, "invalid user address")
		return
	}

	err = h.service.Update(userAddressId, userId, dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.UpdateResponse(ctx, true)
}
func (h *UserAddressHandler) Destroy(ctx *gin.Context) {
	userId, err := h.getAuthUser(ctx)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	idParam := ctx.Param("id")
	userAddressId, err := strconv.Atoi(idParam)
	if err != nil {
		response.ErrorResponse(ctx, "invalid user address")
		return
	}

	err = h.service.Destroy(userAddressId, userId)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.DeletedResponse(ctx)
}

func (h *UserAddressHandler) getAuthUser(ctx *gin.Context) (int, error) {
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
