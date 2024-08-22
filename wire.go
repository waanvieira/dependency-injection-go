//go:build wireinject
// +build wireinject

// A linhas acima são Tags obrigatórias da própria linguagem e lib. são anotações
package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/waanvieira/di-go/product"
)

// Fazemos a inversão de dependencia, mudando a interface para a nossa classe concreta
// toda vez que alguem chamar a nossa interface vai instanciar a nossa classe concreta, no caso o productrepository
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	// Aqui informamos, toda vez que encontrat um  ProductRepositoryInterface vai trocar pelo
	// ProductRepository, e sempre mudamos a classe concreta, assim temos controle na dependencia
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

// Essa função pode ter o nome que quiser
// db - esse parametro é a injeção do DB no nosso repositório, que está sendi indicado no nosso main
// seria basicamente isso  usecase := NewUseCase(db)
// Injetar diretamente em cada struct o db quando precisa, formando assim uma "arvore", nesse caso
// iriamos injetar o repositorio no use case
func NewUseCase(db *sql.DB) *product.ProductUseCase {
	// Aqui nos informamos o build, nesse build nós vinculamos a dependencia com a sua classe
	// Aqui substituimos por exemplo no usecase ter que injetar isso no main toda hora
	// Aqui nós fazemos a injeção de dependencia, parecido com o laravel quando fazemos o bind
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
