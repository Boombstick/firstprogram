package services

import (
	"context"
	"firstprogram/models"
	"fmt"
	"strings"
)

type UserService struct {
	userRepo UserRepository
}

func NewPostgresService(repository UserRepository) *UserService {
	return &UserService{userRepo: repository}
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
		return 0, fmt.Errorf("ошибка вставки пользователя: %w", err)
	}

	return user.Id, nil
}
