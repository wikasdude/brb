package repository

import (
	"brb-midsvc-platform/internal/domain"
	"database/sql"
	"errors"
)

type VendorRepository interface {
	Create(vendor *domain.Vendor) error
	Update(vendor *domain.Vendor) error
	GetByID(id string) (*domain.Vendor, error) // ID is now a string
	ListAll() ([]domain.Vendor, error)
}

type vendorRepository struct {
	db *sql.DB
}

func NewVendorRepository(db *sql.DB) VendorRepository {
	return &vendorRepository{db: db}
}

func (r *vendorRepository) Create(vendor *domain.Vendor) error {
	query := `INSERT INTO vendors (id, name, email, phone) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, vendor.ID, vendor.Name, vendor.Email, vendor.Phone)
	return err
}

func (r *vendorRepository) Update(vendor *domain.Vendor) error {
	query := `UPDATE vendors SET name = $1, email = $2, phone = $3 WHERE id = $4`
	_, err := r.db.Exec(query, vendor.Name, vendor.Email, vendor.Phone, vendor.ID)
	return err
}

func (r *vendorRepository) GetByID(id string) (*domain.Vendor, error) { // ID is now a string
	query := `SELECT id, name, email, phone FROM vendors WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var vendor domain.Vendor
	err := row.Scan(&vendor.ID, &vendor.Name, &vendor.Email, &vendor.Phone)
	if err == sql.ErrNoRows {
		return nil, errors.New("vendor not found")
	}
	return &vendor, err
}

func (r *vendorRepository) ListAll() ([]domain.Vendor, error) {
	query := `SELECT id, name, email, phone FROM vendors`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vendors []domain.Vendor
	for rows.Next() {
		var v domain.Vendor
		if err := rows.Scan(&v.ID, &v.Name, &v.Email, &v.Phone); err != nil {
			return nil, err
		}
		vendors = append(vendors, v)
	}
	return vendors, nil
}
