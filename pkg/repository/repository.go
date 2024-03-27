package repository

import (
	"marketplace/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Item interface {
}

type Repository struct {
	Authorization
	Item
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
