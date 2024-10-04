package rest

import (
	"fmt"
	"go-boiler-plate/config"
	"go-boiler-plate/internal/database"
	"go-boiler-plate/internal/middleware"
	"go-boiler-plate/modules/health"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitializeApiRestServer(cfg *config.Config) {
	router := gin.Default()

	initializeMiddleware(router)
	initializeApiRoutes(router)

	fmt.Println("Http server up and running")
	zap.L().Debug("HTTP Api Server starting : " + fmt.Sprint(cfg.ServerConfig.Port))
	router.Run(fmt.Sprintf(":%v", cfg.ServerConfig.Port))
}

func initializeMiddleware(router *gin.Engine) {
	router.Use(middleware.ErrorMiddleware())
	router.Use()
}

func initializeApiRoutes(router *gin.Engine) {
	var healthFuncs []health.HealthFunc
	healthFuncs = append(healthFuncs, database.Health())
	health.InitHealthRoute(router.Group("/health"), healthFuncs...)
}
