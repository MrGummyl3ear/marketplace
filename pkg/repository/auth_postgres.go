package repository

import (
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
		tx.Rollback()
		return 0, err
	}
	return user.Id, tx.Commit().Error
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	tx := r.db.Begin()
	req := tx.Where("username = ? and password = ?", username, password).First(&user)
	if req.Error != nil {
		tx.Rollback()
		return user, req.Error
	}
	return user, tx.Commit().Error
}
