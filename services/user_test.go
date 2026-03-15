package services

import (
	"context"
	"errors"
	"firstprogram/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type mockUserRepo struct {
	createFunc func(ctx context.Context, user *models.User) error
}

func (m *mockUserRepo) Create(ctx context.Context, user *models.User) error {
	return m.createFunc(ctx, user)
}

func TestPostgresService_CreateUser(t *testing.T) {
	tests := []struct {
		name     string
		userName string
		userAge  int
		mockID   int64
		mockErr  error
		wantID   int64
		wantErr  string
	}{
		{
			name:     "успешное создание",
			userName: "Alex",
			userAge:  21,
			mockID:   1,
			wantID:   1,
		},
		{
			name:     "второй пользователь",
			userName: "Bob",
			userAge:  30,
			mockID:   2,
			wantID:   2,
		},
		{
			name:     "пустое имя",
			userName: "",
			userAge:  21,
			wantErr:  "name не может быть пустым",
		},
		{
			name:     "пробелы в имени",
			userName: "   ",
			userAge:  21,
			wantErr:  "name не может быть пустым",
		},
		{
			name:     "возраст 0",
			userName: "Alex",
			userAge:  0,
			wantErr:  "age должен быть больше 0",
		},
		{
			name:     "отрицательный возраст",
			userName: "Alex",
			userAge:  -5,
			wantErr:  "age должен быть больше 0",
		},
		{
			name:     "ошибка базы",
			userName: "Alex",
			userAge:  21,
			mockErr:  errors.New("connection refused"),
			wantErr:  "ошибка вставки пользователя",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockUserRepo{
				createFunc: func(ctx context.Context, user *models.User) error {
					assert.Equal(t, tt.userName, user.Name)
					assert.Equal(t, tt.userAge, user.Age)

					if tt.mockErr == nil {
						user.Id = tt.mockID
					}
					return tt.mockErr
				},
			}

			svc := NewPostgresService(mock, zap.NewNop())
			id, err := svc.CreateUser(context.Background(), tt.userName, tt.userAge)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Equal(t, int64(0), id)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantID, id)
			}
		})
	}
}
