package middlewares

import (
	"github.com/alimarzban99/ecommerce/config"
	"github.com/alimarzban99/ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var errMessage string
				env := config.Cfg.App.Environment

				if env == "production" {
					errMessage = "خطای داخلی سرور رخ داد. لطفاً بعداً دوباره تلاش کنید."
				} else {
					switch v := r.(type) {
					case string:
						errMessage = v
					case error:
						errMessage = v.Error()
					default:
						errMessage = "خطای ناشناخته‌ای رخ داد"
					}
				}

				response.PanicResponse(ctx, errMessage)
			}
		}()
		ctx.Next()
	}
}
