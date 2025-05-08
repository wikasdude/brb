package repository

import (
	"brb-midsvc-platform/internal/domain"
	"database/sql"
	"time"
)

type BookingRepository interface {
	Create(booking *domain.Booking) error
	GetByID(id int64) (*domain.Booking, error)
	GetVendorBookings(vendorID int64) ([]domain.Booking, error)
	CountByStatus(vendorID int64, status string) (int, error)
	CountTotalByVendor(vendorID int64) (int, error)
	FindOverlappingBooking(vendorID int64, start time.Time, end time.Time) (bool, error)
}

type bookingRepo struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepo{db: db}
}

func (r *bookingRepo) Create(booking *domain.Booking) error {
	query := `
		INSERT INTO bookings (user_id, vendor_id, service_id, start_time, end_time, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		booking.UserID,
		booking.VendorID,
		booking.ServiceID,
		booking.StartTime,
		booking.EndTime,
		booking.Status,
	).Scan(&booking.ID)

	return err
}

func (r *bookingRepo) GetByID(id int64) (*domain.Booking, error) {
	query := `SELECT id, user_id, vendor_id, service_id, start_time, end_time, status, created_at FROM bookings WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var b domain.Booking
	err := row.Scan(&b.ID, &b.UserID, &b.VendorID, &b.ServiceID, &b.StartTime, &b.EndTime, &b.Status, &b.StartTime)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &b, err
}

func (r *bookingRepo) GetVendorBookings(vendorID int64) ([]domain.Booking, error) {
	query := `SELECT id, user_id, vendor_id, service_id, start_time, end_time, status, created_at FROM bookings WHERE vendor_id = $1`
	rows, err := r.db.Query(query, vendorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var b domain.Booking
		err := rows.Scan(&b.ID, &b.UserID, &b.VendorID, &b.ServiceID, &b.StartTime, &b.EndTime, &b.Status, &b.StartTime)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *bookingRepo) CountByStatus(vendorID int64, status string) (int, error) {
	query := `SELECT COUNT(*) FROM bookings WHERE vendor_id = $1 AND status = $2`
	var count int
	err := r.db.QueryRow(query, vendorID, status).Scan(&count)
	return count, err
}

func (r *bookingRepo) CountTotalByVendor(vendorID int64) (int, error) {
	query := `SELECT COUNT(*) FROM bookings WHERE vendor_id = $1`
	var count int
	err := r.db.QueryRow(query, vendorID).Scan(&count)
	return count, err
}

func (r *bookingRepo) FindOverlappingBooking(vendorID int64, start time.Time, end time.Time) (bool, error) {
	query := `
		SELECT COUNT(*) FROM bookings 
		WHERE vendor_id = $1 AND 
		      ((start_time < $3 AND end_time > $2))
	`
	var count int
	err := r.db.QueryRow(query, vendorID, start, end).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
