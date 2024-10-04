package redis

import (
	"context"
	"fmt"
	"go-boiler-plate/config"
	"testing"
	"time"
)

func TestRedisInit(t *testing.T) {
	fmt.Print("Cache 1")
	config.Load("config-test.yaml")
	InitializeRedis(&config.RedisConfig{})

	ctx := context.Background()
	SetKey(&ctx, "one", "one", time.Hour)
}
