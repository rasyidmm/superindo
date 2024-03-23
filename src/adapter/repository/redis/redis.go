package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDataHandler struct {
	redis *redis.Client
}

func NewRedisDataHandler(r *redis.Client) *RedisDataHandler {
	return &RedisDataHandler{
		redis: r,
	}
}

func (r *RedisDataHandler) GetCache(ctx context.Context, key string) (string, error) {
	val, err := r.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "data not found", err
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}
func (r *RedisDataHandler) SetCache(ctx context.Context, key string, value string, duration time.Duration) error {
	return r.redis.Set(ctx, key, value, duration).Err()
}
func (r *RedisDataHandler) DeleteCache(ctx context.Context, key ...string) error {
	return r.redis.Del(ctx, key...).Err()
}
