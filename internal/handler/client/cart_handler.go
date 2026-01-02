package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	resourceClient "github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"

	"github.com/alimarzban99/ecommerce/internal/service/client"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service *client.CartService
}

func NewCartHandler() *CartHandler {
	return &CartHandler{service: client.NewCartService()}
}

func (h *CartHandler) Add(ctx *gin.Context) {
	dto := new(dtoClient.CartAddDTO)
	err := ctx.ShouldBind(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource := user.(*resourceClient.UserResource)
	err = h.service.Add(ctx.Request.Context(), uint(userResource.ID), *dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.CreatedResponse(ctx, nil)
}
func (h *CartHandler) Remove(ctx *gin.Context) {
	dto := new(dtoClient.CartAddDTO)
	err := ctx.ShouldBind(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}

	user, _ := ctx.Get("user")
	userResource := user.(*resourceClient.UserResource)
	err = h.service.Remove(ctx.Request.Context(), uint(userResource.ID), *dto)

	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}

	response.UpdateResponse(ctx, nil)
}
func (h *CartHandler) Finalize(ctx *gin.Context) {

}
