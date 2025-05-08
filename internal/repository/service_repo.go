package repository

import (
	"brb-midsvc-platform/internal/domain"
	"database/sql"
	"errors"
)

type ServiceRepository interface {
	Create(service *domain.Service) error
	Update(service *domain.Service) error
	GetByID(id int64) (*domain.Service, error)
	ListAll() ([]domain.Service, error)
}

type serviceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{db: db}
}

func (r *serviceRepository) Create(service *domain.Service) error {
	query := `INSERT INTO services (name, description, active) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, service.Name, service.Description, service.Active).Scan(&service.ID)
}

func (r *serviceRepository) Update(service *domain.Service) error {
	query := `UPDATE services SET name = $1, description = $2, active = $3 WHERE id = $4`
	_, err := r.db.Exec(query, service.Name, service.Description, service.Active, service.ID)
	return err
}

func (r *serviceRepository) GetByID(id int64) (*domain.Service, error) {
	query := `SELECT id, name, description, active FROM services WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var service domain.Service
	err := row.Scan(&service.ID, &service.Name, &service.Description, &service.Active)
	if err == sql.ErrNoRows {
		return nil, errors.New("service not found")
	}
	return &service, err
}

func (r *serviceRepository) ListAll() ([]domain.Service, error) {
	query := `SELECT id, name, description, active FROM services`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []domain.Service
	for rows.Next() {
		var s domain.Service
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Active); err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}
