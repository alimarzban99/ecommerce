package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/internal/service/client"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *client.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: client.NewUserService()}
}

func (h *UserHandler) Profile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.AuthenticationErrorResponse(c, "User not found in context")
		return
	}

	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.PanicResponse(c, "Invalid user type in context")
		return
	}

	response.SuccessResponse(c, userResource)
}

func (h *UserHandler) Update(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		response.AuthenticationErrorResponse(c, "User not found in context")
		return
	}

	userResource, ok := user.(*resourceClient.UserResource)
	if !ok {
		response.PanicResponse(c, "Invalid user type in context")
		return
	}

	dto := new(dtoClient.UpdateProfileDTO)
	if err := c.ShouldBindJSON(dto); err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(c, validationErrors)
		return
	}

	updatedUser, err := h.service.UpdateProfile(userResource.ID, dto)
	if err != nil {
		response.ErrorResponse(c, err.Error())
		return
	}

	response.UpdateResponse(c, updatedUser)
}
