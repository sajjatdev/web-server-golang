package service

import "github.com/sajjatdev/invoice_management/internal/domain"

type CustomerService interface {
	CreateCustomer(customer *domain.Customer) error
	GetCustomer(id string) (*domain.Customer, error)
	GetAllCustomers() ([]*domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) error
	DeleteCustomer(id string) error
}

type ProductService interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id string) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id string) error
}

type InvoiceService interface {
	CreateInvoice(invoice *domain.Invoice) error
	GetInvoice(id string) (*domain.Invoice, error)
	GetAllInvoices() ([]*domain.Invoice, error)
	UpdateInvoice(invoice *domain.Invoice) error
	DeleteInvoice(id string) error
	AddLineItem(invoiceID string, line *domain.InvoiceLine) error
}
