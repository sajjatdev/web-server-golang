package service

import (
	"time"

	"github.com/sajjatdev/invoice_management/internal/domain"
	"github.com/sajjatdev/invoice_management/internal/repository"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{customerRepo: customerRepo}
}

func (s *customerService) CreateCustomer(customer *domain.Customer) error {
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Now()
	return s.customerRepo.Create(customer)
}

func (s *customerService) GetCustomer(id string) (*domain.Customer, error) {
	return s.customerRepo.GetByID(id)
}

func (s *customerService) GetAllCustomers() ([]*domain.Customer, error) {
	return s.customerRepo.GetAll()
}

func (s *customerService) UpdateCustomer(customer *domain.Customer) error {
	customer.UpdatedAt = time.Now()
	return s.customerRepo.Update(customer)
}

func (s *customerService) DeleteCustomer(id string) error {
	return s.customerRepo.Delete(id)
}
