package repositories

import (
	"firstprogram/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type UserRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateTable() error {
	return r.db.Model(&models.User{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
}

func (r *UserRepository) Create(user *models.User) error {
	_, err := r.db.Model(user).Insert()
	return err
}
