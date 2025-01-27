package service

import (
	"InventoryManagementSystem/models"
	"InventoryManagementSystem/repository"
)

type InventoryService struct {
	inventoryRepo  repository.InventoryRepository
	productService ProductService
}

func NewInventoryService(inventoryRepo repository.InventoryRepository, service ProductService) *InventoryService {
	return &InventoryService{inventoryRepo, service}
}

func (i *InventoryService) AddInventory(productId int, quantity int) error {
	_, err := i.productService.GetProductById(productId)
	if err != nil {
		return err
	}
	return i.inventoryRepo.AddInventory(productId, quantity)
}

func (i *InventoryService) RemoveInventory(productId int) error {
	_, err := i.productService.GetProductById(productId)
	if err != nil {
		return err
	}
	return i.inventoryRepo.RemoveInventory(productId, 1)
}

func (i *InventoryService) RemoveBulkInventory(productId int, quantity int) error {
	_, err := i.productService.GetProductById(productId)
	if err != nil {
		return err
	}
	return i.inventoryRepo.RemoveInventory(productId, quantity)
}

func (i *InventoryService) ViewInventory() ([]models.Inventory, error) {
	return i.inventoryRepo.ViewInventory()
}

func (i *InventoryService) GetInventoryById(productId int) (*models.Inventory, error) {
	return i.inventoryRepo.GetInventoryById(productId)
}
