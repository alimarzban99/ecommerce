package repository

import (
	"github.com/alimarzban99/ecommerce/internal/dto/auth"
	"github.com/alimarzban99/ecommerce/internal/model"
	authResources "github.com/alimarzban99/ecommerce/internal/resources/auth"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"time"
)

type VerificationCodeRepository struct {
	*Repository[model.VerificationCode, auth.CreateOTPCodeDTO, auth.CreateOTPCodeDTO, authResources.CodeResponse]
}

func NewVerificationCodeRepository() *VerificationCodeRepository {
	return &VerificationCodeRepository{
		&Repository[model.VerificationCode, auth.CreateOTPCodeDTO, auth.CreateOTPCodeDTO, authResources.CodeResponse]{
			database: database.DB(),
		},
	}
}

func (r VerificationCodeRepository) ValidCode(dto *auth.VerifyOTPCodeDTO) (bool, error) {
	var exists bool

	err := r.database.
		Model(&model.VerificationCode{}).
		Select("count(*) > 0").
		Where("mobile = ?", dto.Mobile).
		Where("code = ?", dto.Code).
		Where("expire_at >= ?", time.Now()).
		Find(&exists).
		Error

	return exists, err
}
