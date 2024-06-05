package types

import (
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uuid.UUID) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() (*[]Product, error)
	GetProductByID(productID uuid.UUID) (*Product, error)
	GetProductsByID(productIDs *[]string) (*[]Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image"`
	Price       float64   `json:"price" db:"price"`
	Quantity    int       `json:"quantity" db:"quantity"` // not atomic, can cause issues with multiple concurrent requests
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}
