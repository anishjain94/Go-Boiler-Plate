package rest

import (
	"fmt"
	"go-boiler-plate/infra/environment"
	"go-boiler-plate/infra/middleware"
	"go-boiler-plate/modules/health"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitializeApiRestServer() {
	router := gin.Default()

	initializeMiddleware(router)
	initializeApiRoutes(router)

	fmt.Println("Http server up and running")
	zap.L().Debug("HTTP Api Server starting : " + fmt.Sprint(environment.PORT))
	router.Run(fmt.Sprintf(":%v", environment.PORT))
}

func initializeMiddleware(router *gin.Engine) {
	router.Use(middleware.ErrorMiddleware())
	router.Use()
}

func initializeApiRoutes(router *gin.Engine) {
	health.InitHealthRoute(router.Group("/health"))
}
