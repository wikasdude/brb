package main

import (
	"brb-midsvc-platform/api/handler"
	"brb-midsvc-platform/internal/repository"
	"brb-midsvc-platform/internal/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to Postgres
	db, err := sql.Open("postgres", "postgres://booking:booking@localhost:5432/brb?sslmode=disable")
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}
	defer db.Close()

	log.Println("✅ Connected to PostgreSQL")

	// Initialize Repository
	bookingRepo := repository.NewBookingRepository(db)

	// Initialize Usecase
	bookingUC := usecase.NewBookingUsecase(bookingRepo)

	// Initialize Handler
	h := handler.NewHandler(bookingUC)

	// Register Routes
	http.HandleFunc("/health", h.HealthCheck)
	http.HandleFunc("/summary/vendor/", h.VendorSummary)
	http.HandleFunc("/bookings", h.CreateBooking)

	log.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
