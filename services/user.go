package services

import (
	"context"
	"firstprogram/models"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type UserService struct {
	userRepo UserRepository
	logger   *zap.Logger
}

func NewPostgresService(repository UserRepository, logger *zap.Logger) *UserService {
	return &UserService{userRepo: repository,
		logger: logger.Named("user_service")}
}

func (s *UserService) CreateUser(ctx context.Context, name string, age int) (int64, error) {

	if strings.TrimSpace(name) == "" {
		return 0, &ValidationError{Field: "Name", Message: "name не может быть пустым"}
	}
	if age <= 0 {
		return 0, &ValidationError{Field: "Age", Message: "age должен быть больше 0"}
	}
	user := &models.User{
		Name: name,
		Age:  age,
	}
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		s.logger.Error("ошибка создания пользователя",
			zap.String("name", name),
			zap.Int("age", age),
			zap.Error(err))
		return 0, fmt.Errorf("ошибка создания пользователя: %w", err)
	}
	s.logger.Info("пользователь создан",
		zap.Int64("id", user.Id),
		zap.String("name", name),
	)
	return user.Id, nil
}
