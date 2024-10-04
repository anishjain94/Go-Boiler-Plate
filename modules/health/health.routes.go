package health

import (
	"go-boiler-plate/internal/ratelimiter"
	"go-boiler-plate/util"

	"github.com/gin-gonic/gin"
)

func InitHealthRoute(router *gin.RouterGroup, healthFuncs ...HealthFunc) {
	healthRoute(router, healthFuncs...)
}

func healthRoute(router *gin.RouterGroup, healthFunc ...HealthFunc) {
	h := Health{HealthFuncs: healthFunc}
	router.GET("", ratelimiter.Limit(ratelimiter.RateLimit30PerMinute), util.HandleHTTPGet(h.getHealth))
}
