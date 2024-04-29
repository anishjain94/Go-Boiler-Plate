package redis

import (
	"context"
	"fmt"
	"go-boiler-plate/infra/environment"
	"testing"
	"time"
)

func TestRedisInit(t *testing.T) {
	fmt.Print("Cache 1")
	environment.InitializeEnvs()

	InitializeRedis()

	ctx := context.Background()
	SetKey(&ctx, "one", "one", time.Hour)

}
