package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func (p Product) GetPriceWithTax(taxRate float64) (float64, error) {
	if taxRate < 0 { // simple err handling
		return 0, errors.New("tax rate cannot be negative")
	}

	tax := p.Price * taxRate
	return p.Price + tax, nil
}

func main() {
	// if many field can directly assign prod := Product{...}
	var product Product
	product.Name = "Laptop"
	product.Price = 1200.0

	priceWithTax, err := product.GetPriceWithTax(0.07)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Product: %s\n", product.Name)
	fmt.Printf("Price with tax: %.2f\n", priceWithTax)
}
