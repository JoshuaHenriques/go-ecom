package user

// Repository

import (
	"fmt"

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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := &types.User{}
	err := s.db.Get(user, "SELECT * FROM user WHERE email=$1", email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *Store) GetUserByID(id uuid.UUID) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}

/*
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := &types.User

	err := rows.Scan(
		user.ID,
		user,FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
*/
