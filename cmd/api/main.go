package main

import (
	"fmt"
	"github.com/alimarzban99/ecommerce/config"
	"github.com/alimarzban99/ecommerce/internal/container"
	"github.com/alimarzban99/ecommerce/internal/middlewares"
	"github.com/alimarzban99/ecommerce/internal/routers"
	"github.com/alimarzban99/ecommerce/pkg/cache"
	"github.com/alimarzban99/ecommerce/pkg/database"
	"github.com/alimarzban99/ecommerce/pkg/jwt"
	logger "github.com/alimarzban99/ecommerce/pkg/logging"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.Load()

	if err := database.Init(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := cache.Init(); err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	zapLogger, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	defer zapLogger.Sync()

	// Load JWT keys once at startup
	privateKey, publicKey, err := jwt.LoadKeys("keys/private.pem", "keys/public.pem")
	if err != nil {
		log.Fatalf("Failed to load JWT keys: %v", err)
	}

	// Create dependency injection container
	c := container.NewContainer(database.DB(), privateKey, publicKey)

	appConfig := config.Cfg.App
	gin.SetMode(appConfig.Environment)
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(middlewares.CustomRecovery())
	router.Use(middlewares.Throttle())
	router.Use(middlewares.ZapLogger(zapLogger))

	routers.HealthRouter(router.Group(""))

	apiV1 := router.Group("api/v1/")
	routers.AuthRouter(apiV1, c)
	routers.UserRouter(apiV1, c)
	routers.CategoryRouter(apiV1, c)
	routers.ProductRouter(apiV1, c)
	routers.OrderRouter(apiV1, c)
	routers.CartRouter(apiV1, c)

	runPort := fmt.Sprintf(":%d", config.Cfg.Server.Port)
	log.Printf("Server is running at http://localhost%s\n", runPort)
	err = router.Run(runPort)
	if err != nil {
		log.Fatal(err)
	}
}
