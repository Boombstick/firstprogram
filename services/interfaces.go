package services

import (
	"context"

	"firstprogram/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

type Cache interface {
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
}

type IUserService interface {
	CreateUser(ctx context.Context, name string, age int) (int64, error)
}

type ICounterService interface {
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
}
