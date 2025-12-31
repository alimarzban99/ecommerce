package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 {
			log.Error("http_request",
				append(fields, zap.String("errors", c.Errors.String()))...,
			)
			return
		}

		switch c.Writer.Status() {
		case 400, 401, 403, 404:
			log.Warn("http_request", fields...)
		case 500:
			log.Error("http_request", fields...)
		default:
			log.Info("http_request", fields...)
		}
	}
}
