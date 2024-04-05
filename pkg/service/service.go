package service

import (
	"marketplace/pkg/model"
	"marketplace/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Item interface {
	Create(item model.Item) error
	GetAllItems(params model.QueryParam) ([]model.Item,uint64, error)
}

type Service struct {
	Authorization
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Item:          NewItemService(repos.Item),
	}
}
