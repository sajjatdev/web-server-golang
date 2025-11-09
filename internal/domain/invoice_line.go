package domain

type InvoiceLine struct {
	ID        string   `json:"id" db:"id"`
	InvoiceID string   `json:"invoice_id" db:"invoice_id"`
	ProductID string   `json:"product_id" db:"product_id"`
	Quantity  int      `json:"quantity" db:"quantity"`
	UnitPrice float64  `json:"unit_price" db:"unit_price"`
	Total     float64  `json:"total" db:"total"`
	Product   *Product `json:"product,omitempty" db:"-"`
}
