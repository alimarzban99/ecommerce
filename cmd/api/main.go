package main

import (
	"fmt"
	"github.com/alimarzban99/ecommerce/config"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/alimarzban99/ecommerce/internal/model"
	"github.com/alimarzban99/ecommerce/internal/routers"
	"github.com/alimarzban99/ecommerce/pkg/cache"
	"github.com/alimarzban99/ecommerce/pkg/database"
	logger "github.com/alimarzban99/ecommerce/pkg/logging"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Load()

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	if err := cache.Init(); err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	zapLogger, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	model.Starter()

	appConfig := config.Cfg.App
	gin.SetMode(appConfig.Environment)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middlewares.CustomRecovery())
	router.Use(middlewares.Throttle())
	router.Use(middlewares.ZapLogger(zapLogger))

	routers.HealthRouter(router.Group(""))

	apiV1 := router.Group("api/v1/")
	routers.AuthRouter(apiV1)
	routers.UserRouter(apiV1)
	routers.CategoryRouter(apiV1)
	routers.ProductRouter(apiV1)
	routers.OrderRouter(apiV1)

	runPort := fmt.Sprintf(":%d", config.Cfg.Server.Port)
	log.Printf("Server is running at http://localhost%s\n", runPort)
	err = router.Run(runPort)
	if err != nil {
		log.Fatal(err)
	}
}
