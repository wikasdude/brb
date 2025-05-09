package handler

import (
	"brb-midsvc-platform/internal/domain"
	"brb-midsvc-platform/internal/usecase"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	DB             *sql.DB
	BookingUsecase usecase.BookingUsecase
}

func NewHandler(bookingUC usecase.BookingUsecase) *Handler {
	return &Handler{BookingUsecase: bookingUC}
}
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/summary/vendor/", h.VendorSummary)
}

// HealthCheck godoc
// @Summary Check service health
// @Description Checks if the database connection is alive
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /health [get]
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.DB.Ping(); err != nil {
		http.Error(w, `{"status":"DB disconnected"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"OK"}`))
}

// VendorSummary godoc
// @Summary Get summary for a vendor
// @Description Returns booking summary for a given vendor
// @Tags Vendor
// @Accept  json
// @Produce  json
// @Param   vendor_id  path  int  true  "Vendor ID"
// @Success 200 {object} domain.VendorSummary
// @Failure 400 {object} map[string]string
// @Router /summary/vendor/{vendor_id} [get]
func (h *Handler) VendorSummary(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Missing vendor ID", http.StatusBadRequest)
		return
	}
	vendorIDStr := parts[3]
	vendorID, err := strconv.ParseInt(vendorIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid vendor ID", http.StatusBadRequest)
		return
	}

	summary, err := h.BookingUsecase.GetVendorSummary(vendorID)
	if err != nil {
		http.Error(w, "Failed to fetch summary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Creates a booking with the given payload
// @Tags booking
// @Accept json
// @Produce json
// @Param booking body domain.Booking true "Booking Payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /bookings [post]
func (h *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking *domain.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	fmt.Println("Booking:", booking.EndTime)

	err := h.BookingUsecase.CreateBooking(booking)
	if err != nil {
		http.Error(w, "Failed to create booking: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Booking ID:", booking.ID)
	fmt.Println("booking created")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Booking created",
		"id":      booking.ID,
	})
}
