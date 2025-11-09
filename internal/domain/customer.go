package domain

import (
	"time"
)

type Customer struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CustomerRepository interface {
	Create(customer *Customer) error
	GetByID(id string) (*Customer, error)
	GetAll() ([]*Customer, error)
	Update(customer *Customer) error
	Delete(id string) error
}
