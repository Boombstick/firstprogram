package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(host, port string) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis: %w", err)
	}

	fmt.Println("Redis: подключено")
	return &RedisCache{client: client}, nil
}

func (r *RedisCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.client.IncrBy(ctx, key, value).Result()
}
