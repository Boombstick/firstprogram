package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewPostgres(cfg PgConfig) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database})

	if _, err := db.Exec("SELECT 1"); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к PostgreSQL: %w", err)
	}
	return db, nil

}
