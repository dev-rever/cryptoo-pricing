package mredis

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/dev-rever/cryptoo-pricing/config"
	api "github.com/dev-rever/cryptoo-pricing/utils/apiutils"
	logger "github.com/dev-rever/cryptoo-pricing/utils/logutils"
)

type Wrap struct {
	redis *redis.Client
}

func ProvideMRedis() *Wrap {

	opt := redis.Options{
		Addr: config.GetRedisAddr(),
	}
	red := redis.NewClient(&opt)
	return &Wrap{redis: red}
}

func (w *Wrap) RateLimitMiddleware(maxAttempts int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		key := "login_attempts:" + ip

		count, err := w.redis.Incr(ctx, key).Result()
		if err != nil {
			logger.LogError(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
			return
		}

		if count == 1 {
			w.redis.Expire(ctx, key, window)
		}

		if count > int64(maxAttempts) {
			err := errors.New("too many requests")
			logger.LogError(err)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, api.ResponseError(api.InternalErrorCode, err.Error()))
			return
		}

		c.Next()
	}
}
