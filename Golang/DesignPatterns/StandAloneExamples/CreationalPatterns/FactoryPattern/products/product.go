package products

import "time"

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func (p *Product) NewProduct(prodName string) *Product {
	product := Product{
		ProductName: prodName,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	return &product
}
