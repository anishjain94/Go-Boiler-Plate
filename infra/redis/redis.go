package redis

import (
	"context"
	"go-boiler-plate/util"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const MSG_REDIS_INIT_ERROR = "Redis Connection Failed"

var (
	redisClient *redis.Client
)

func InitializeRedis() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Username: os.Getenv("REDIS_USER"),
		DB:       int(1),
	})

	pong, err := redisClient.Ping(context.Background()).Result()
	util.AssertError(err, http.StatusInternalServerError, "REDIS INTI ERROR")

	zap.L().Info("Redis Successfully Connected : " + redisClient.String() + " : " + pong)

}

func Info(ctx *context.Context) (string, error) {
	result, err := redisClient.Info(*ctx).Result()

	return result, err
}

func SetKey(ctx *context.Context, key string, value string, expiry time.Duration) error {

	err := redisClient.Set(*ctx, key, value, expiry).Err()

	return err
}

func GetKey(ctx *context.Context, key string) (string, error) {
	val, err := redisClient.Get(*ctx, key).Result()

	return val, err
}

func DelKey(ctx *context.Context, key string) error {
	err := redisClient.Del(*ctx, key).Err()

	return err
}

func GetIdleTime(ctx *context.Context, key string) int {
	idleTime := redisClient.ObjectIdleTime(*ctx, key)
	return int(idleTime.Val().Seconds())
}

func Keys(ctx *context.Context, pattern string) []string {
	keys, err := redisClient.Keys(*ctx, pattern).Result()
	if err != nil {
		return []string{}
	}

	return keys
}

func Incr(ctx *context.Context, key string) int64 {
	val, err := redisClient.Incr(*ctx, key).Result()
	if err != nil {
		return 0
	}

	return val
}

func SetExpiry(ctx *context.Context, key string, expiry time.Duration) bool {
	val, err := redisClient.Expire(*ctx, key, expiry).Result()
	if err != nil {
		return false
	}
	return val
}

func AddToSet(ctx *context.Context, key string, value string) bool {

	count, err := redisClient.SAdd(*ctx, key, value).Result()
	if err != nil {
		return false
	}

	return count == 1
}

func RemoveFromSet(ctx *context.Context, key string, value string) bool {

	count, err := redisClient.SRem(*ctx, key, value).Result()
	if err != nil {
		return false
	}

	return count == 1
}

func LengthOfSet(ctx *context.Context, key string) int64 {

	count, err := redisClient.SCard(*ctx, key).Result()
	if err != nil {
		return 0
	}

	return count
}

func InSet(ctx *context.Context, key string, value string) bool {
	member, err := redisClient.SIsMember(*ctx, key, value).Result()
	if err != nil {
		return false
	}

	return member
}

func GetALLSetValues(ctx *context.Context, key string) []string {
	members, err := redisClient.SMembers(*ctx, key).Result()
	if err != nil {
		return nil
	}

	return members
}

func RemoveAllKeys(ctx *context.Context, keysKey string, deleteCondition func(key string) bool) {
	conditionalDelete := deleteCondition != nil

	allKeys := GetALLSetValues(ctx, keysKey)

	for _, key := range allKeys {
		DelKey(ctx, key)
		if conditionalDelete && deleteCondition(key) {
			RemoveFromSet(ctx, keysKey, key)
		}
	}

	if !conditionalDelete {
		DelKey(ctx, keysKey)
	}
}

func SetKeyMap(ctx *context.Context, keysKey string, key string, expiry time.Duration) {
	count := LengthOfSet(ctx, keysKey)
	added := AddToSet(ctx, keysKey, key)

	firstEntry := added && count == 0
	if firstEntry {
		SetExpiry(ctx, keysKey, expiry)
	}
}

func IsNilErr(err error) bool {
	return err == redis.Nil
}
