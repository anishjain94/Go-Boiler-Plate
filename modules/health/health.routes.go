package health

import (
	"go-boiler-plate/util"

	"github.com/gin-gonic/gin"
)

func InitHealthRoute(router *gin.RouterGroup) {
	healthRoute(router)
}

func healthRoute(router *gin.RouterGroup) {
	router.GET("/", util.HandleHTTPGet(getHealth))
}
