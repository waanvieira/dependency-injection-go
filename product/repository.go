package product

import "database/sql"

// Aqui indicamos que vamos receber uma interface para poder fazer a inversão de dependencia
type ProductRepositoryInterface interface {
	GetProduct(id int) (*Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) GetProduct(id int) (*Product, error) {
	return &Product{
		ID:   id,
		Name: "Product Name",
	}, nil
}
