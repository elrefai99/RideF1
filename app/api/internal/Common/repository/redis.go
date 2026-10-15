package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func Set(ctx context.Context, client *redis.Client, key string, value any, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	return client.Get(ctx, key).Result()
}
