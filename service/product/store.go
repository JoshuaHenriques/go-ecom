package product

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
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

func (s *Store) GetProductByID(productID uuid.UUID) (*types.Product, error) {
	product := &types.Product{}
	err := s.db.Get(product, "SELECT * FROM products WHERE id=$1", productID.String())
	if err != nil {
		return nil, fmt.Errorf("product not found: %v", err)
	}

	return product, nil
}

func (s *Store) GetProductsByID(productIDs *[]string) (*[]types.Product, error) {
	products := &[]types.Product{}
	query := fmt.Sprintf(`SELECT * FROM products WHERE id IN ('%s'::uuid)`, strings.Join(*productIDs, "'::uuid,'"))
	err := s.db.Select(products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	_, err := s.db.Exec("INSERT INTO products (name, price, image, description, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Price, product.Image, product.Description, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}
