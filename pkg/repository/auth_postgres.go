package repository

import (
	"log"
	"marketplace/pkg/model"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {

	tx := r.db.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	return 0, tx.Commit().Error
}
