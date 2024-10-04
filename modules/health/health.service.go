package health

import (
	"github.com/gin-gonic/gin"
)

func getHealth(c *gin.Context) *HealthResponseDto {
	return &HealthResponseDto{
		Message: "Healthy",
	}
}
