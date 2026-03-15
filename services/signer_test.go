package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignHMACSHA512(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		key      string
		expected string
		wantErr  string
	}{
		{
			name:     "успешная подпись",
			text:     "test",
			key:      "test123",
			expected: "b596e24739fd44d42ffd25f26ea367dad3a71f61c8c5fab6b6ee6ceeae5a7170b66445d6eaadfb49e6d4e968a2888726ff522e3bf065c966aa66a24153778382",
		},
		{
			name:    "пустой text",
			text:    "",
			key:     "test123",
			wantErr: "text не может быть пустым",
		},
		{
			name:    "пробелы в text",
			text:    "   ",
			key:     "test123",
			wantErr: "text не может быть пустым",
		},
		{
			name:    "пустой key",
			text:    "test",
			key:     "",
			wantErr: "key не может быть пустым",
		},
		{
			name:    "пробелы в key",
			text:    "test",
			key:     "   ",
			wantErr: "key не может быть пустым",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SignHMACSHA512(tt.text, tt.key)

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
				assert.Len(t, result, 128)
			}
		})
	}
}

func TestSignHMACSHA512_Deterministic(t *testing.T) {
	sig1, _ := SignHMACSHA512("test", "key")
	sig2, _ := SignHMACSHA512("test", "key")
	assert.Equal(t, sig1, sig2, "одинаковые входы должны давать одинаковый результат")
}

func TestSignHMACSHA512_DifferentInputs(t *testing.T) {
	sig1, _ := SignHMACSHA512("hello", "key")
	sig2, _ := SignHMACSHA512("world", "key")
	assert.NotEqual(t, sig1, sig2, "разные тексты должны давать разные подписи")

	sig3, _ := SignHMACSHA512("test", "key1")
	sig4, _ := SignHMACSHA512("test", "key2")
	assert.NotEqual(t, sig3, sig4, "разные ключи должны давать разные подписи")
}
