package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type mockCache struct {
	result int64
	err    error
}

func (m *mockCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return m.result, m.err
}

func TestRedisService_IncrBy(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      int64
		mockResult int64
		mockErr    error
		wantResult int64
		wantErr    string
	}{
		{
			name:       "успешный инкремент",
			key:        "age",
			value:      19,
			mockResult: 19,
			wantResult: 19,
		},
		{
			name:       "повторный инкремент",
			key:        "age",
			value:      5,
			mockResult: 24,
			wantResult: 24,
		},
		{
			name:       "отрицательное значение",
			key:        "age",
			value:      -3,
			mockResult: 21,
			wantResult: 21,
		},
		{
			name:    "пустой key",
			key:     "",
			value:   10,
			wantErr: "key не может быть пустым",
		},
		{
			name:    "пробелы в key",
			key:     "   ",
			value:   10,
			wantErr: "key не может быть пустым",
		},
		{
			name:    "ошибка Redis",
			key:     "age",
			value:   10,
			mockErr: errors.New("connection refused"),
			wantErr: "ошибка Redis INCRBY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockCache{result: tt.mockResult, err: tt.mockErr}
			svc := NewRedisService(mock, zap.NewNop())

			result, err := svc.IncrBy(context.Background(), tt.key, tt.value)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResult, result)
			}
		})
	}
}
