package domain

import (
	"time"
)

type Product struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	ImageData   []byte    `json:"image_data" db:"image_data"`
	SKU         string    `json:"sku" db:"sku"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id string) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id string) error
}
