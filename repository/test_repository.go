package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/p-glynn/okeefe-ecg-v2-backend/models"
)

type TestRepository struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) *TestRepository {
	return &TestRepository{db: db}
}

func (r *TestRepository) Create(test *models.Test) error {
	query := `
		INSERT INTO tests (user_id, title, description, ecg_data, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(
		query,
		test.UserID,
		test.Title,
		test.Description,
		test.ECGData,
		test.Status,
	).Scan(&test.ID, &test.CreatedAt, &test.UpdatedAt)
}

func (r *TestRepository) GetByID(id int64) (*models.Test, error) {
	test := &models.Test{}
	query := `
		SELECT id, user_id, title, description, ecg_data, status, created_at, updated_at
		FROM tests
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&test.ID,
		&test.UserID,
		&test.Title,
		&test.Description,
		&test.ECGData,
		&test.Status,
		&test.CreatedAt,
		&test.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("test not found")
	}
	return test, err
}

func (r *TestRepository) GetByUserID(userID int64) ([]*models.Test, error) {
	query := `
		SELECT id, user_id, title, description, ecg_data, status, created_at, updated_at
		FROM tests
		WHERE user_id = $1
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []*models.Test
	for rows.Next() {
		test := &models.Test{}
		err := rows.Scan(
			&test.ID,
			&test.UserID,
			&test.Title,
			&test.Description,
			&test.ECGData,
			&test.Status,
			&test.CreatedAt,
			&test.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}
	return tests, rows.Err()
}

func (r *TestRepository) Update(test *models.Test) error {
	query := `
		UPDATE tests
		SET title = $1, description = $2, ecg_data = $3, status = $4, updated_at = $5
		WHERE id = $6
		RETURNING updated_at`

	return r.db.QueryRow(
		query,
		test.Title,
		test.Description,
		test.ECGData,
		test.Status,
		time.Now(),
		test.ID,
	).Scan(&test.UpdatedAt)
}
