package auth

import (
	authdto "github.com/alimarzban99/ecommerce/internal/dto/auth"
	"github.com/alimarzban99/ecommerce/internal/service/auth"

	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/alimarzban99/ecommerce/pkg/validation"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *auth.Service
}

func NewAuthHandler() *Handler {
	return &Handler{service: auth.NewAuthService()}
}

func (h *Handler) GetVerificationCode(ctx *gin.Context) {
	dto := new(authdto.GetOTPCodeDTO)
	err := ctx.ShouldBindJSON(&dto)
	if err != nil {
		validationErrors := validation.FormatValidationErrors(err)
		response.ValidationErrorsResponse(ctx, validationErrors)
		return
	}
	h.service.GetVerificationCode(dto)
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
	jti := ctx.MustGet("jti").(string)
	h.service.Logout(jti)
	response.SuccessResponse(ctx, "Logged out")
}
