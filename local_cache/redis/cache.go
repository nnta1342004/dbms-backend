package rediscache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hareta/appCommon"
	localcache "hareta/local_cache"
	"time"
)

type localCache struct {
	rdc *redis.Client
}

func NewLocalCache(rdc *redis.Client) localcache.LocalCache {
	return &localCache{rdc: rdc}
}
func (s *localCache) SetNX(ctx context.Context, key string, data interface{}, ttl time.Duration) error {
	str, err := appCommon.MarshalData(data)
	if err != nil {
		return err
	}
	return s.rdc.SetNX(ctx, key, str, ttl).Err()
}
func (s *localCache) Get(ctx context.Context, key string, data interface{}) error {
	val, err := s.rdc.Get(ctx, key).Result()
	if err == nil {
		if err := appCommon.StringToJson(val, data); err != nil {
			return err
		}
		return nil
	}
	return err
}
