package middlewares

import (
	"schedule_gateway/global"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimiter(baseKey string, rate int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, exists := c.Get("user_id")
		if !exists {
			c.Next()
			return
		}

		rdb := global.RedisDb.Client
		ctx := c.Request.Context()
		key := baseKey + "_" + userId.(string)

		// Try to initialize (first request)
		created, err := rdb.SetNX(ctx, key, rate-1, duration).Result()
		if err != nil {
			c.Next() // fail open
			return
		}

		if created {
			c.Next()
			return
		}

		// Atomic decrement (TTL preserved)
		remaining, err := rdb.Decr(ctx, key).Result()
		if err != nil {
			c.Next() // or abort, depending on policy
			return
		}

		if remaining < 0 {
			c.AbortWithStatus(429)
			return
		}

		c.Next()
	}
}
