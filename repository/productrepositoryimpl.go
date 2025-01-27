package repository

import (
	"InventoryManagementSystem/models"
	"errors"
)

type ProductRepositoryImpl struct {
	products map[int]*models.Product
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		products: make(map[int]*models.Product),
	}
}

func (p *ProductRepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	p.products[product.Id] = product
	return product, nil
}

func (p *ProductRepositoryImpl) GetProductById(id int) (*models.Product, error) {
	val, ok := p.products[id]
	if ok {
		return val, nil
	}
	return nil, errors.New("product not found")
}

func (p *ProductRepositoryImpl) UpdateProduct(id int, price int) (*models.Product, error) {
	product, ok := p.products[id]
	if ok {
		product.Price = price
		return product, nil
	}
	return nil, errors.New("product not found")
}

func (p *ProductRepositoryImpl) DeleteProduct(id int) error {
	_, ok := p.products[id]
	if ok {
		delete(p.products, id)
		return nil
	}
	return errors.New("product not found")
}
