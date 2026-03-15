package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	client *redis.Client
	logger *zap.Logger
}

func NewRedisCache(host, port string, logger *zap.Logger) (*Redis, error) {
	log := logger.Named("redis")
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis: %w", err)
	}
	log.Info("подключено",
		zap.String("host", host),
		zap.String("port", port))
	return &Redis{client: client}, nil
}

func (r *Redis) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return r.client.IncrBy(ctx, key, value).Result()
}
