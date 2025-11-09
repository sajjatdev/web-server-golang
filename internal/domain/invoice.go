package domain

import (
	"time"
)

type InvoiceStatus string

const (
	InvoiceStatusDraft     InvoiceStatus = "draft"
	InvoiceStatusSent      InvoiceStatus = "sent"
	InvoiceStatusPaid      InvoiceStatus = "paid"
	InvoiceStatusOverdue   InvoiceStatus = "overdue"
	InvoiceStatusCancelled InvoiceStatus = "cancelled"
)

type Invoice struct {
	ID            string        `json:"id" db:"id"`
	CustomerID    string        `json:"customer_id" db:"customer_id"`
	InvoiceNumber string        `json:"invoice_number" db:"invoice_number"`
	IssueDate     time.Time     `json:"issue_date" db:"issue_date"`
	DueDate       time.Time     `json:"due_date" db:"due_date"`
	Status        InvoiceStatus `json:"status" db:"status"`
	SubTotal      float64       `json:"sub_total" db:"sub_total"`
	TaxAmount     float64       `json:"tax_amount" db:"tax_amount"`
	TotalAmount   float64       `json:"total_amount" db:"total_amount"`
	Notes         string        `json:"notes" db:"notes"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`

	Customer  *Customer      `json:"customer,omitempty" db:"-"`
	LineItems []*InvoiceLine `json:"line_items,omitempty" db:"-"`
}

type InvoiceRepository interface {
	Create(invoice *Invoice) error
	GetByID(id string) (*Invoice, error)
	GetAll() ([]*Invoice, error)
	GetByCustomerID(customerID string) ([]*Invoice, error)
	Update(invoice *Invoice) error
	Delete(id string) error
	AddLineItem(line *InvoiceLine) error
	GetLineItems(invoiceID string) ([]*InvoiceLine, error)
	RemoveLineItem(lineID string) error
}
