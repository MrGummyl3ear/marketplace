package repository

import (
	"marketplace/pkg/model"

	"gorm.io/gorm"
)

type ItemPostgres struct {
	db *gorm.DB
}

func NewItemPostgres(db *gorm.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(item model.Item) error {
	{

		tx := r.db.Begin()
		err := tx.Omit("created_at").Create(&item).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		return tx.Commit().Error
	}
}

func (r *ItemPostgres) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	tx := r.db.Begin()
	err := tx.Find(&items).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return items, tx.Commit().Error
}
