package repository

import (
	"brb-midsvc-platform/internal/domain"
	"database/sql"
	"errors"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	ListAll() ([]domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email, role, password) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, user.Name, user.Email, user.Role, user.Password).Scan(&user.ID)
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	query := `SELECT id, name, email, role, password FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	query := `SELECT id, name, email, role, password FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *userRepository) ListAll() ([]domain.User, error) {
	query := `SELECT id, name, email, role, password FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
