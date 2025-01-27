package service

import (
	"InventoryManagementSystem/models"
	"InventoryManagementSystem/repository"
	"InventoryManagementSystem/util"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p *ProductService) CreateProduct(name string, price int, category string) (*models.Product, error) {
	product := &models.Product{}
	product.Name = name
	product.Price = price
	product.Category = models.ProductCategory(category)
	product.Id = util.GenerateID()
	return p.productRepo.CreateProduct(product)
}
func (p *ProductService) GetProductById(id int) (*models.Product, error) {
	return p.productRepo.GetProductById(id)
}

func (p *ProductService) UpdateProduct(id int, price int) (*models.Product, error) {
	return p.productRepo.UpdateProduct(id, price)
}
func (p *ProductService) DeleteProduct(id int) error {
	return p.productRepo.DeleteProduct(id)
}
