package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/p-glynn/okeefe-ecg-v2-backend/models"
)

type CommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	query := `
		INSERT INTO comments (test_id, user_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(
		query,
		comment.TestID,
		comment.UserID,
		comment.Content,
	).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
}

func (r *CommentRepository) GetByID(id int64) (*models.Comment, error) {
	comment := &models.Comment{}
	query := `
		SELECT id, test_id, user_id, content, created_at, updated_at
		FROM comments
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.TestID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("comment not found")
	}
	return comment, err
}

func (r *CommentRepository) GetByTestID(testID int64) ([]*models.Comment, error) {
	query := `
		SELECT id, test_id, user_id, content, created_at, updated_at
		FROM comments
		WHERE test_id = $1
		ORDER BY created_at ASC`

	rows, err := r.db.Query(query, testID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		comment := &models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.TestID,
			&comment.UserID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, rows.Err()
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	query := `
		UPDATE comments
		SET content = $1, updated_at = $2
		WHERE id = $3
		RETURNING updated_at`

	return r.db.QueryRow(
		query,
		comment.Content,
		time.Now(),
		comment.ID,
	).Scan(&comment.UpdatedAt)
}
