package auth

import (
	authdto "github.com/alimarzban99/ecommerce/internal/dto/auth"
	"github.com/alimarzban99/ecommerce/internal/service"

	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetVerificationCode(ctx *gin.Context) {
	dto := new(authdto.GetOTPCodeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}
	err = h.service.GetVerificationCode(dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}
	response.SuccessResponse(ctx, "Code sent")
}

func (h *Handler) Verify(ctx *gin.Context) {
	dto := new(authdto.VerifyOTPCodeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}
	token, err := h.service.Verify(dto)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}
	response.SuccessResponse(ctx, struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	})
}

func (h *Handler) Logout(ctx *gin.Context) {
	jtiValue, exists := ctx.Get("jti")
	if !exists {
		response.AuthenticationErrorResponse(ctx, "Token ID not found")
		return
	}
	jti, ok := jtiValue.(string)
	if !ok {
		response.ErrorResponse(ctx, "Invalid token ID type")
		return
	}
	err := h.service.Logout(jti)
	if err != nil {
		response.ErrorResponse(ctx, err.Error())
		return
	}
	response.SuccessResponse(ctx, "Logged out")
}
