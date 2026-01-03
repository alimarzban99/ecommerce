package middlewares

import (
	"crypto/rsa"
	"fmt"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// Authentication creates a middleware that validates JWT tokens
// It accepts publicKey, tokenRepo, and userRepo as dependencies
func Authentication(
	publicKey *rsa.PublicKey,
	tokenRepo repository.TokenRepositoryInterface,
	userRepo repository.UserRepositoryInterface,
	kind string,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 8 {
			response.AuthenticationErrorResponse(ctx, "Authentication required")
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil || !token.Valid {
			response.AuthenticationErrorResponse(ctx, "Unauthorized")
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.AuthenticationErrorResponse(ctx, "invalid token")
			ctx.Abort()
			return
		}

		// Safe type assertion for jti
		jtiValue, ok := claims["jti"]
		if !ok {
			response.AuthenticationErrorResponse(ctx, "invalid token: missing jti")
			ctx.Abort()
			return
		}
		jti, ok := jtiValue.(string)
		if !ok {
			response.AuthenticationErrorResponse(ctx, "invalid token: invalid jti type")
			ctx.Abort()
			return
		}

		// Verify token exists and is not revoked
		tokenExists, err := tokenRepo.FindToken(jti)
		if err != nil || !tokenExists {
			response.AuthenticationErrorResponse(ctx, "invalid token")
			ctx.Abort()
			return
		}

		// Safe type assertion for user ID
		subValue, ok := claims["sub"]
		if !ok {
			response.AuthenticationErrorResponse(ctx, "invalid token: missing sub")
			ctx.Abort()
			return
		}

		var userID int
		switch v := subValue.(type) {
		case float64:
			userID = int(v)
		case int:
			userID = v
		case int64:
			userID = int(v)
		default:
			response.AuthenticationErrorResponse(ctx, "invalid token: invalid sub type")
			ctx.Abort()
			return
		}

		user, err := userRepo.FindOne(userID)
		if err != nil {
			response.AuthenticationErrorResponse(ctx, "user not found")
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Set("jti", jti)
		ctx.Next()
	}
}
