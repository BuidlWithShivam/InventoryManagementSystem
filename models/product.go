package models

type ProductCategory string

const (
	Electronics ProductCategory = "Electronics"
	Grocery                     = "Grocery"
	Clothing                    = "Clothing"
	Furniture                   = "Furniture"
)

type Product struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Price    int             `json:"price"`
	Category ProductCategory `json:"category"`
}
