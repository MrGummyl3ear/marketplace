package repository

import (
	"marketplace/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Item interface {
	Create(item model.Item) error
	GetAllItems(params model.QueryParam) ([]model.Item,uint64,  error)
}

type Repository struct {
	Authorization
	Item
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Item:          NewItemPostgres(db),
	}
}
