package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	RedisPort        string `mapstructure:"REDIS_PORT"`
	RedisHost        string `mapstructure:"REDIS_HOST"`
}

func Load() (*Config, error) {

	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Ошибка чтения .env: %w", err)

	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфига: %w", err)
	}

	return &cfg, nil
}
