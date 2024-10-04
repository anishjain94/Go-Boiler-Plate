package ratelimiter

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// Init multiple rate-limits to be used on different routes
const (
	RateLimit2PerMinute    = "2_per_minute"
	RateLimit10PerMinute   = "10_per_minute"
	RateLimit30PerMinute   = "30_per_minute"
	RateLimit60PerMinute   = "60_per_minute"
	RateLimit120PerMinute  = "120_per_minute"
	RateLimit300PerMinute  = "300_per_minute"
	RateLimit1000PerMinute = "1000_per_minute"
	RateLimit5000PerMinute = "5000_per_minute"
	RateLimit100PerDay     = "100_per_day"
	RateLimit10PerDay      = "10_per_day"
)

// { name, secondsDuration, allowedRequests }
var rateLimitConfigs = []RateLimitConfig{
	{RateLimit2PerMinute, 60, 2},
	{RateLimit10PerMinute, 60, 10},
	{RateLimit30PerMinute, 60, 30},
	{RateLimit100PerDay, 86400, 100},
	{RateLimit10PerDay, 86400, 10},
	{RateLimit60PerMinute, 60, 60},
	{RateLimit120PerMinute, 60, 120},
	{RateLimit300PerMinute, 60, 300},
	{RateLimit1000PerMinute, 60, 1000},
	{RateLimit5000PerMinute, 60, 5000},
}

// RateLimitConfig defines the configuration for rate limiting
type RateLimitConfig struct {
	Name            string
	SecondsDuration int64
	AllowedRequests int64
}

/*
Limit is a middleware that rate limits based on IP address.
  - @secondDuration(default 60): The time period in seconds for which the request limit is applied
  - @allowedRequests(default 120): The number of requests allowed in the given time period
*/
func Limit(rateLimiterName string) gin.HandlerFunc {
	var cfg *RateLimitConfig
	for _, conf := range rateLimitConfigs {
		if conf.Name == rateLimiterName {
			cfg = &conf
			break
		}
	}

	if cfg == nil {
		slog.Error(fmt.Sprintf("No rate limiter found with name: %s", rateLimiterName))
		os.Exit(1)
	}

	// Init rate limiter
	rate := limiter.Rate{
		Period: time.Duration(cfg.SecondsDuration) * time.Second,
		Limit:  cfg.AllowedRequests,
	}
	store := memory.NewStore()
	ipRateLimiter := limiter.New(store, rate)

	return func(c *gin.Context) {
		identifier := c.ClientIP()

		limiterCtx, err := ipRateLimiter.Get(c.Request.Context(), identifier)
		if err != nil {
			slog.Error(fmt.Sprintf("IPRateLimit - ipRateLimiter.Get - err: %v, %s on %s", err, identifier, c.Request.URL))
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			c.Abort() // Stop further processing
			return
		}

		// Set rate limit headers
		h := c.Writer.Header()
		h.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
		h.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
		h.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

		if limiterCtx.Reached {
			slog.Error(fmt.Sprintf("Too Many Requests from %s on %s", identifier, c.Request.URL))
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too Many Requests on " + c.Request.URL.String(),
			})
			c.Abort() // Stop further processing
			return
		}

		c.Next() // Proceed to the next middleware/handler
	}
}
