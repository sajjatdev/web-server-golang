package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sajjatdev/invoice_management/internal/handler"
	"github.com/sajjatdev/invoice_management/internal/repository"
	"github.com/sajjatdev/invoice_management/internal/service"
	"github.com/sajjatdev/invoice_management/pkg/database"
)

func main() {
	// Initialize database
	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	customerRepo := repository.NewCustomerRepository(db)
	// productRepo := repository.NewProductRepository(db)
	// invoiceRepo := repository.NewInvoiceRepository(db)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)
	// productService := service.NewProductService(productRepo)
	// invoiceService := service.NewInvoiceService(invoiceRepo, productRepo)

	// Initialize router
	r := mux.NewRouter()

	// Setup routes
	handler.SetupRoutes(r, customerService)

	// Get server port from environment
	port := os.Getenv("SERVER_PORT")

	// Start server
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
