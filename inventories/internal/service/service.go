package service

import (
	"context"
	model "mpg/inventories/internal/model"
	repo "mpg/inventories/internal/repository"
)

type InventorySvcInterface interface {
	Create(ctx context.Context, inventoryPreliminary *model.Inventory) (*model.Inventory, error)
	GetById(ctx context.Context, id string) (*model.Inventory, error)
	Update(ctx context.Context, id string, updateFields *model.Inventory) (int, error)
	Delete(ctx context.Context, id string) (int, error)
}

type InventoryService struct {
	repo repo.InventoryRepository
}


func NewInventoryService(repo repo.InventoryRepository) InventorySvcInterface {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) Create(ctx context.Context, inventoryPreliminary *model.Inventory) (*model.Inventory, error) {
	inventoryPreliminary.ItemsID = make([]string, 0)
	result, err := s.repo.Create(ctx, inventoryPreliminary)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *InventoryService) GetById(ctx context.Context, id string) (*model.Inventory, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *InventoryService) Update(ctx context.Context, id string, updateFields *model.Inventory) (int, error) {
	result, err := s.repo.Update(ctx, id, *updateFields)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *InventoryService) Delete(ctx context.Context, id string) (int, error) {
	result, err := s.repo.Delete(ctx, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}