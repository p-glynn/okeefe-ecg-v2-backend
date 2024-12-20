package repository

import (
	"database/sql"
	"errors"
	"time"

	// "github.com/lib/pq"
	"server/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (email, password_hash, first_name, last_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, first_name, last_name, created_at, updated_at
		FROM users
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, password_hash, first_name, last_name, created_at, updated_at
		FROM users
		WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, first_name = $2, last_name = $3, updated_at = $4
		WHERE id = $5
		RETURNING updated_at`

	return r.db.QueryRow(
		query,
		user.Email,
		user.FirstName,
		user.LastName,
		time.Now(),
		user.ID,
	).Scan(&user.UpdatedAt)
}
