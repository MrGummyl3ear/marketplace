package model

import "time"

type Item struct {
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_post;sort:desc"` //подправить время
	Price     int       `json:"price" binding:"required" gorm:"index:idx_post,"`
	Title     string    `json:"title" binding:"required" gorm:"not null;size:256"`
	Content   string    `json:"content"  gorm:"size:256"`
	Username  string    `json:"username" gorm:"unique not null"`
}

//добавить изображения
