package order

import (
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

func (s *Store) CreateOrder(order types.Order) (uuid.UUID, error) {
	var id uuid.UUID
	err := s.db.QueryRowx("INSERT INTO orders (user_id, total, status, address) VALUES ($1, $2, $3, $4) returning id", order.UserID, order.Total, order.Status, order.Address).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, err
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
