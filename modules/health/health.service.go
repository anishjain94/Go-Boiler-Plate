package health

import (
	"context"
)

func getHealth(ctx *context.Context) *HealthResponseDto {

	return &HealthResponseDto{
		Message: "Healthy",
	}
}
