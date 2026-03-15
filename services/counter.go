package services

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type CounterService struct {
	cache  Cache
	logger *zap.Logger
}

func NewRedisService(cache Cache, logger *zap.Logger) *CounterService {
	return &CounterService{cache: cache,
		logger: logger.Named("counter_service")}
}

func (s *CounterService) IncrBy(ctx context.Context, key string, value int64) (int64, error) {

	if strings.TrimSpace(key) == "" {
		return 0, &ValidationError{Field: "Key", Message: "key не может быть пустым"}
	}

	result, err := s.cache.IncrBy(ctx, key, value)
	if err != nil {
		s.logger.Error("ошибка инкремента",
			zap.String("key", key),
			zap.Int64("value", value),
			zap.Error(err),
		)
		return 0, fmt.Errorf("ошибка Redis INCRBY: %w", err)
	}
	s.logger.Info("инкремент выполнен",
		zap.String("key", key),
		zap.Int64("value", value),
		zap.Int64("result", result),
	)
	return result, nil
}
