package repository

import (
	"fmt"
	"marketplace/pkg/model"
	"math"

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

func (r *ItemPostgres) GetAllItems(params model.QueryParam) ([]model.Item, uint64, error) {
	var items []model.Item
	tx := r.db.Begin()
	offset := (params.Page - 1) * params.PageSize
	err := tx.Scopes(OrderByField(params), FilterPrice(params), Pagination(int(offset), int(params.PageSize))).Find(&items).Error
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	var total int64
	tx.Model(&model.Item{}).Scopes(FilterPrice(params)).Count(&total)
	MaxPage := uint64(math.Ceil(float64(total) / float64(params.PageSize)))
	return items, MaxPage, tx.Commit().Error
}

func FilterPrice(params model.QueryParam) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("price >= %d and price <= %d", params.MinPrice, params.MaxPrice))
	}
}

func OrderByField(params model.QueryParam) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", params.Order, params.Sort))
	}
}

func Pagination(offset int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
