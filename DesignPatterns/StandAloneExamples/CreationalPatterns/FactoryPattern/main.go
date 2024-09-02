package main

import (
	"factory/products"
	"fmt"
)

func main() {
	factory := products.Product{}
	product := factory.NewProduct("Cooling Machine")
	fmt.Println("My Product", product.ProductName, "is created at,", product.CreatedAt.UTC())
}
