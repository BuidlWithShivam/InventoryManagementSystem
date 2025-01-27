package main

import (
	"InventoryManagementSystem/repository"
	"InventoryManagementSystem/service"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	productService := service.NewProductService(repository.NewProductRepositoryImpl())
	inventoryService := service.NewInventoryService(repository.NewInventoryRepositoryImpl(), *productService)

	product, err := productService.CreateProduct("Samsung Mobile", 50000, "Electronics")
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	err = inventoryService.AddInventory(product.Id, 100)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	err = inventoryService.RemoveInventory(product.Id)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	inventory, err := inventoryService.GetInventoryById(product.Id)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println(*inventory)

	err = inventoryService.RemoveBulkInventory(product.Id, 5)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}

	inventory, err = inventoryService.GetInventoryById(product.Id)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println(*inventory)

	inventories, err := inventoryService.ViewInventory()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return
	}
	fmt.Println(inventories)
}

// Create two models Product and Inventory
// Create Product Service as well : add/remove/update product
// Repository interface which can be extended to multiple data access
