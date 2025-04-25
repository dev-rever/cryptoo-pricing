package redisutil

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	. "github.com/dev-rever/cryptoo-pricing/utils"
)

type MRedis struct {
	cli *redis.Client
}

func ProvideRedis() *MRedis {
	opt := redis.Options{
		Addr: "localhost:6379",
	}
	red := redis.NewClient(&opt)
	return &MRedis{cli: red}
}

func (rd *MRedis) RateLimitMiddleware(maxAttempts int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()
		key := "login_attempts:" + ip

		count, err := rd.cli.Incr(ctx, key).Result()
		if err != nil {
			LogError(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(InternalErrorCode, err.Error()))
			return
		}

		if count == 1 {
			rd.cli.Expire(ctx, key, window)
		}

		if count > int64(maxAttempts) {
			err := errors.New("too many requests")
			LogError(err)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, ResponseError(UnknownErrorCode, err.Error()))
			return
		}

		c.Next()
	}
}
