package redis

import (
	"context"
	"time"
)

type RedisRepository interface {
	GetCache(ctx context.Context, key string) (string, error)
	SetCache(ctx context.Context, key string, value string, duration time.Duration) error
	DeleteCache(ctx context.Context, key ...string) error
}
