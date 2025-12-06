package order

import "fmt"

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func (p *Product) GetDetails() string {
	return "Product ID: " + p.ID + ", Name: " + p.Name + ", Price: $" + fmt.Sprintf("%.2f", p.Price)
}
