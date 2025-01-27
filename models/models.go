package models

import (
	"encoding/json"

	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	PasswordHash  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Test struct {
	ID          int64           `json:"id"`
	UserID      int64           `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ECGData     json.RawMessage `json:"ecg_data"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type Comment struct {
	ID        int64     `json:"id"`
	TestID    int64     `json:"test_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
