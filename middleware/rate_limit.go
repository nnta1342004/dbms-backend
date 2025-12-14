package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	goservice "github.com/leductoan3082004/go-sdk"
	"hareta/appCommon"
	"time"
)

const rateLimit = 500

func RateLimit(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		rdc := sc.MustGet(appCommon.PluginRedis).(*redis.Client)
		ip := c.ClientIP()
		key := fmt.Sprintf("LIMIT_RATE_IP_%s", ip)
		//rdc.Del(c.Request.Context(), key)
		err := rdc.Watch(c.Request.Context(), func(tx *redis.Tx) error {
			_ = rdc.SetNX(c.Request.Context(), key, 0, time.Minute)

			count, err := rdc.Incr(c.Request.Context(), key).Result()
			if count > rateLimit {
				panic(appCommon.ErrInvalidRequest(errors.New("You have requested so many")))
			}
			if err != nil {
				return err
			}
			return nil
		}, key)
		if err != nil {
			panic(appCommon.ErrInternal(err))
		}
		c.Next()
	}
}
