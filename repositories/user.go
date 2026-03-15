package repositories

import (
	"context"
	"firstprogram/models"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go.uber.org/zap"
)

type UserRepository struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewUserRepository(ctx context.Context, db *pg.DB, logger *zap.Logger) (*UserRepository, error) {

	log := logger.Named("user_repository")
	repo := UserRepository{db: db}
	err := repo.initDatabaseTables(ctx)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания таблицы: %w", err)
	}
	log.Info("таблица users готова")
	return &repo, nil
}

func (r *UserRepository) initDatabaseTables(ctx context.Context) error {
	return r.db.WithContext(ctx).Model(&models.User{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.WithContext(ctx).Model(user).Insert()
	return err
}
