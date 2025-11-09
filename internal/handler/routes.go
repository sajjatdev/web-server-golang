package handler

import (
	"github.com/gorilla/mux"
	"github.com/sajjatdev/invoice_management/internal/service"
)

func SetupRoutes(
	r *mux.Router,
	customerService service.CustomerService,
	// productService service.ProductService,
	// invoiceService service.InvoiceService,
) {
	customerHandler := NewCustomerHandler(customerService)
	// productHandler := NewProductHandler(productService)
	// invoiceHandler := NewInvoiceHandler(invoiceService)

	// Customer routes
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", customerHandler.GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")

	// // Product routes
	// r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	// r.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	// r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	// r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	// r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// // Invoice routes
	// r.HandleFunc("/invoices", invoiceHandler.CreateInvoice).Methods("POST")
	// r.HandleFunc("/invoices", invoiceHandler.GetAllInvoices).Methods("GET")
	// r.HandleFunc("/invoices/{id}", invoiceHandler.GetInvoice).Methods("GET")
	// r.HandleFunc("/invoices/{id}", invoiceHandler.UpdateInvoice).Methods("PUT")
	// r.HandleFunc("/invoices/{id}", invoiceHandler.DeleteInvoice).Methods("DELETE")
	// r.HandleFunc("/invoices/{id}/lines", invoiceHandler.AddLineItem).Methods("POST")
}
