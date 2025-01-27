package repository

import "InventoryManagementSystem/models"

type InventoryRepository interface {
	AddInventory(productId int, quantity int) error
	RemoveInventory(productId int, quantity int) error
	ViewInventory() ([]models.Inventory, error)
	GetInventoryById(productId int) (*models.Inventory, error)
}
