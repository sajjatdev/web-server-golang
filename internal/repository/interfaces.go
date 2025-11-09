package repository

import "github.com/sajjatdev/invoice_management/internal/domain"

type CustomerRepository interface {
	domain.CustomerRepository
}

type ProductRepository interface {
	domain.ProductRepository
}

type InvoiceRepository interface {
	domain.InvoiceRepository
}
