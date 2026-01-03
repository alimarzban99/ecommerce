package routers

import (
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/handler/auth"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup, c *container.Container) {
	authRouter := r.Group("auth")

	authHandler := auth.NewAuthHandler(c.AuthService)
	authRouter.POST("get-verification-code", authHandler.GetVerificationCode)
	authRouter.POST("verify", authHandler.Verify)
	authRouter.GET("logout",
		middlewares.Authentication(
			c.PublicKey,
			c.TokenRepository,
			c.UserRepository,
			"admin",
		),
		authHandler.Logout,
	)
}
