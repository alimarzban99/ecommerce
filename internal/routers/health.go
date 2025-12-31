package routers

import (
	"github.com/alimarzban99/ecommerce/internal/handler/health"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	healthHandler := health.NewHealthHandler()

	r.GET("health", healthHandler.Health)
}
