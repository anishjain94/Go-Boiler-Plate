package health

import (
	"context"
	util "go-boiler-plate/util"
	"net/http"
)

func getHealth(ctx *context.Context) *HealthResponseDto {

	util.ErrorIf(true, http.StatusBadGateway, "unknown")

	return &HealthResponseDto{
		Message: "Healthy",
	}
}
