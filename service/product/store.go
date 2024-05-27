package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joshuahenriques/go-ecom/types"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() (*[]types.Product, error) {
	products := &[]types.Product{}
	err := s.db.Select(products, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("products not found: %v", err)
	}

	return products, nil
}
