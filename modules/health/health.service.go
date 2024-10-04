package health

import (
	"github.com/gin-gonic/gin"
)

type HealthFunc func() (name string, status string)

type Health struct {
	HealthFuncs []HealthFunc
}

func (h *Health) getHealth(c *gin.Context) *HealthResponseDto {
	dependencies := make(map[string]string)
	for _, healthFunc := range h.HealthFuncs {
		name, status := healthFunc()
		dependencies[name] = status
	}
	return &HealthResponseDto{
		Message:      "Healthy",
		Dependencies: dependencies,
	}
}
