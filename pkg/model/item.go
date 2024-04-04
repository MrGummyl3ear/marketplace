package model

import "time"

type Item struct {
	CreatedAt time.Time `json:"created_at" gorm:"column:date;index:idx_post;sort:asc"`
	Price     int       `json:"price" binding:"required" gorm:"index:idx_post,"`
	Title     string    `json:"title" binding:"required" gorm:"not null;size:120"`
	Content   string    `json:"content"  gorm:"size:256"`
	Username  string    `json:"username" gorm:"unique not null"`
}

//добавить изображения

type QueryParam struct {
	MaxPrice uint64
	MinPrice uint64
	Page     uint64
	MaxPage  uint64
	PageSize uint64
	NextPage string
	PrevPage string
	Sort     string
	Order    string
}
