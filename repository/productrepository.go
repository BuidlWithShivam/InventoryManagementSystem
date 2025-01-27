package repository

import "InventoryManagementSystem/models"

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductById(id int) (*models.Product, error)
	UpdateProduct(id int, price int) (*models.Product, error)
	DeleteProduct(id int) error
}
