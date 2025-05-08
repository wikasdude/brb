package usecase

import (
	"brb-midsvc-platform/internal/domain"
	"brb-midsvc-platform/internal/repository"
	"errors"
	"fmt"
	"log"
	"time"
)

type BookingUsecase interface {
	CreateBooking(booking *domain.Booking) error
	GetBookingByID(id int64) (*domain.Booking, error)
	GetVendorSummary(vendorID int64) (domain.VendorSummary, error)
}

type bookingUsecase struct {
	repo repository.BookingRepository
}

func NewBookingUsecase(repo repository.BookingRepository) BookingUsecase {
	return &bookingUsecase{repo: repo}
}

// CreateBooking validates and creates a booking
func (uc *bookingUsecase) CreateBooking(booking *domain.Booking) error {

	startTime, err := time.Parse(time.RFC3339, booking.StartTime)
	if err != nil {
		log.Fatalf("Error parsing start time: %v", err)
	}

	endTime, err := time.Parse(time.RFC3339, booking.EndTime)
	if err != nil {
		log.Fatalf("Error parsing end time: %v", err)
	}
	startHour := startTime.Hour()
	endHour := endTime.Hour()

	fmt.Println("Start Time:", startTime)
	fmt.Println("End Time:", endTime)

	fmt.Println("Start Hour:", startHour)
	fmt.Println("End Hour:", endHour)
	if startHour < 9 || endHour > 17 {
		return errors.New("bookings must be between 9 AM and 5 PM and last 1 hour")
	}

	// Check for overlaps
	overlap, err := uc.repo.FindOverlappingBooking(booking.VendorID, startTime, endTime)
	if err != nil {
		return err
	}
	if overlap {
		fmt.Println("Booking overlaps with an existing booking")
		return errors.New("booking time overlaps with an existing booking")
	}

	// Set initial status
	booking.Status = "pending"

	return uc.repo.Create(booking)
}

func (uc *bookingUsecase) GetBookingByID(id int64) (*domain.Booking, error) {
	return uc.repo.GetByID(id)
}

func (uc *bookingUsecase) GetVendorSummary(vendorID int64) (domain.VendorSummary, error) {
	total, err := uc.repo.CountTotalByVendor(vendorID)
	if err != nil {
		return domain.VendorSummary{}, err
	}

	statuses := []string{"pending", "confirmed", "completed"}
	statusCounts := make(map[string]int)

	for _, status := range statuses {
		count, err := uc.repo.CountByStatus(vendorID, status)
		if err != nil {
			return domain.VendorSummary{}, err
		}
		statusCounts[status] = count
	}

	return domain.VendorSummary{
		VendorID:      vendorID,
		TotalBookings: total,
		StatusCounts:  statusCounts,
	}, nil
}
