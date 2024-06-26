package service

import (
	"marketplace/pkg/model"
	"marketplace/pkg/repository"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(item model.Item) error {
	return s.repo.Create(item)
}

func (s *ItemService) GetAllItems(params model.QueryParam) ([]model.Item,uint64,  error) {
	return s.repo.GetAllItems(params)
}
