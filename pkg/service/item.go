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

func (s *ItemService) GetAllItems(params model.QueryParam) ([]model.Item, error) {
	return s.repo.GetAllItems(params)
}

func (r *ItemService) GetMaxPage(params model.QueryParam) (uint64, error) {
	return r.repo.GetMaxPage(params)
}
