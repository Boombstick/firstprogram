package services

import (
	"context"
	"fmt"
	"strings"
)

type CounterService struct {
	cache Cache
}

func NewRedisService(cache Cache) *CounterService {
	return &CounterService{cache: cache}
}

func (s *CounterService) IncrBy(ctx context.Context, key string, value int64) (int64, error) {

	if strings.TrimSpace(key) == "" {
		return 0, &ValidationError{Field: "Key", Message: "key не может быть пустым"}
	}

	result, err := s.cache.IncrBy(ctx, key, value)
	if err != nil {
		return 0, fmt.Errorf("ошибка Redis INCRBY: %w", err)
	}
	return result, nil
}
