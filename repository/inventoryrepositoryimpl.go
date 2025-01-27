package repository

import (
	"InventoryManagementSystem/models"
	"errors"
)

type InventoryRepositoryImpl struct {
	store map[int]int
}

func NewInventoryRepositoryImpl() *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{
		store: make(map[int]int),
	}
}

func (i *InventoryRepositoryImpl) AddInventory(productId int, quantity int) error {
	currQuantity, ok := i.store[productId]
	if ok {
		i.store[productId] = currQuantity + quantity
		return nil
	}
	i.store[productId] = quantity
	return nil
}

func (i *InventoryRepositoryImpl) RemoveInventory(productId int, quantity int) error {
	currentQuantity, ok := i.store[productId]
	if !ok {
		return errors.New("Product not found")
	}
	if currentQuantity < quantity {
		return errors.New("Quantity Not Enough")
	}
	i.store[productId] = currentQuantity - quantity
	return nil
}

func (i *InventoryRepositoryImpl) ViewInventory() ([]models.Inventory, error) {
	var inventory []models.Inventory
	for productId, quantity := range i.store {
		inventory = append(inventory, models.Inventory{
			ProductId: productId,
			Quantity:  quantity,
		})
	}
	return inventory, nil
}

func (i *InventoryRepositoryImpl) GetInventoryById(productId int) (*models.Inventory, error) {
	quantity, ok := i.store[productId]
	if !ok {
		return nil, errors.New("Product not found")
	}
	return &models.Inventory{
		ProductId: productId,
		Quantity:  quantity,
	}, nil
}
