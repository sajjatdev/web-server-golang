package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sajjatdev/invoice_management/internal/domain"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(customer *domain.Customer) error {
	customer.ID = uuid.New().String()
	query := `INSERT INTO customers (id, name, email, phone, address, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(query, customer.ID, customer.Name, customer.Email,
		customer.Phone, customer.Address, customer.CreatedAt, customer.UpdatedAt)
	return err
}

func (r *customerRepository) GetByID(id string) (*domain.Customer, error) {
	query := `SELECT id, name, email, phone, address, created_at, updated_at 
	          FROM customers WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var customer domain.Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone,
		&customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) GetAll() ([]*domain.Customer, error) {
	query := `SELECT id, name, email, phone, address, created_at, updated_at FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*domain.Customer
	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone,
			&customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *customerRepository) Update(customer *domain.Customer) error {
	query := `UPDATE customers SET name=$1, email=$2, phone=$3, address=$4, updated_at=$5 
	          WHERE id=$6`
	_, err := r.db.Exec(query, customer.Name, customer.Email, customer.Phone,
		customer.Address, customer.UpdatedAt, customer.ID)
	return err
}

func (r *customerRepository) Delete(id string) error {
	query := "DELETE FROM customers WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
