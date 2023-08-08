package main

import (
	"fmt"

	"github.com/dwarvesf/go23/ex7/pkg/model"
	"github.com/dwarvesf/go23/ex7/pkg/repo"
)

func main() {

	r := repo.NewPGRepo("postgres://postgres:postgres@localhost:5432/ex7_local?sslmode=disable")

	r.AutoMigrate(&model.Product{}, &model.Vendor{})

	p, err := r.Product().CreateProduct(model.Product{
		Name:     "Product 3",
		SKU:      "SKU-3",
		Price:    100000000.00,
		Currency: "VND",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("created product:", p)

	fmt.Println("Starting server...")

	fmt.Println("Ending server...")
}
