package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"time"
)

type RedisRepository struct {
	mock.Mock
}

func (m *RedisRepository) GetCache(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), nil
}
func (m *RedisRepository) SetCache(ctx context.Context, key string, value string, duration time.Duration) error {

	args := m.Called(ctx, key, value, duration)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (m *RedisRepository) DeleteCache(ctx context.Context, key ...string) error {
	args := m.Called(ctx, key)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
