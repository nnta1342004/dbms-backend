package localcache

import (
	"context"
	"time"
)

type LocalCache interface {
	SetNX(ctx context.Context, key string, data interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, data interface{}) error
}
