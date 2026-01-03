package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/alimarzban99/ecommerce/config"
	"github.com/alimarzban99/ecommerce/internal/dto/auth"
	client2 "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"strconv"
	"time"
)

type Service struct {
	repo       repository.VerificationCodeRepositoryInterface
	tokenRepo  repository.TokenRepositoryInterface
	userRepo   repository.UserRepositoryInterface
	privateKey *rsa.PrivateKey
}

// NewAuthService creates a new auth service (kept for backward compatibility)
func NewAuthService() *Service {
	// This should not be used in production - use NewAuthServiceWithDeps instead
	panic("NewAuthService() is deprecated. Use NewAuthServiceWithDeps() with dependency injection")
}

// NewAuthServiceWithDeps creates a new auth service with injected dependencies
func NewAuthServiceWithDeps(
	repo repository.VerificationCodeRepositoryInterface,
	tokenRepo repository.TokenRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
	privateKey *rsa.PrivateKey,
) *Service {
	return &Service{
		repo:       repo,
		tokenRepo:  tokenRepo,
		userRepo:   userRepo,
		privateKey: privateKey,
	}
}

func (s *Service) GetVerificationCode(dto *auth.GetOTPCodeDTO) error {
	expireTime := config.Cfg.OTPCode.ExpireTime

	_, err := s.repo.Create(&auth.CreateOTPCodeDTO{
		Code:     s.generateOTPCode(),
		Mobile:   dto.Mobile,
		ExpireAt: time.Now().Add(expireTime),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Verify(dto *auth.VerifyOTPCodeDTO) (string, error) {

	var user *client.UserResource
	var err error

	isMobileExists := s.existsByMobile(dto.Mobile)
	if !isMobileExists {
		user, err = s.userRepo.Create(&client2.StoreUserDTO{Mobile: &dto.Mobile})
	} else {
		user, err = s.userRepo.FindByMobile(dto.Mobile)
	}

	if err != nil {
		return "", err
	}
	isCodeValid := s.checkOTPCode(dto)
	if !isCodeValid {
		return "", errors.New("code invalid")
	}

	expiration := time.Now().Add(time.Hour * 24)
	fmt.Println(expiration)
	tokenData, err := s.tokenRepo.Create(&auth.TokenCreate{UserID: uint(user.ID), ExpiresAt: expiration, Revoked: false})

	if err != nil {
		return "", err
	}

	// Get the ID from TokenResponse
	tokenID := tokenData.ID.String()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expiration.Unix(),
		"jti": tokenID,
	})

	tokenStr, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", errors.New("failed to sign token: " + err.Error())
	}

	return tokenStr, nil
}

func (s *Service) Logout(jti string) error {
	return s.tokenRepo.ExpiredToken(jti)
}

func (s *Service) generateOTPCode() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	otp := r.Intn(9000) + 1000
	return strconv.Itoa(otp)
}

func (s *Service) existsByMobile(mobile string) bool {

	exists, err := s.userRepo.CheckExistsByMobile(mobile)

	if err != nil || !exists {
		return false
	}
	return true
}

func (s *Service) checkOTPCode(dto *auth.VerifyOTPCodeDTO) bool {

	ok, err := s.repo.ValidCode(dto)
	if err != nil || !ok {
		return false
	}

	return true
}
