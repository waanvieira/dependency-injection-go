package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// Exemplo de não usar o wire e usar o wire para injetar nossas dependencias
	// Aqui injetamos todos as nossas dependencias a mão no main
	// Create a new product repository
	// repository := product.NewProductRepository(db)
	// Create a new product use case
	// usecase := product.NewProductUseCase(repository)
	// Call method getProduct
	// product, err := usecase.GetProduct(1)

	// Utilizando o wire injetamos apenas o db no NewUseCase
	usecase := NewUseCase(db)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
