package repository

import (
	"github.com/alimarzban99/ecommerce/internal/dto/auth"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/pkg/database"

	authResources "github.com/alimarzban99/ecommerce/internal/resources/auth"
)

type TokenRepository struct {
	*Repository[model.Token, auth.TokenCreate, auth.TokenCreate, authResources.TokenResponse]
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		&Repository[model.Token, auth.TokenCreate, auth.TokenCreate, authResources.TokenResponse]{
			database: database.DB(),
		},
	}
}

func (r *TokenRepository) FindToken(jti string) (bool, error) {
	var exists bool
	err := r.database.
		Model(&model.Token{}).
		Select("count(*) > 0").
		Where("id = ? AND revoked = ?", jti, false).
		Find(&exists).Error
	return exists, err
}

func (r *TokenRepository) ExpiredToken(jti string) error {
	return r.database.
		Model(&model.Token{}).
		Where("id = ?", jti).
		Update("revoked", true).Error
}
