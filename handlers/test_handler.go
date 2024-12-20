package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"server/models"
	"server/repository"
)

type TestHandler struct {
	repo *repository.TestRepository
}

func NewTestHandler(repo *repository.TestRepository) *TestHandler {
	return &TestHandler{repo: repo}
}

type CreateTestRequest struct {
	UserID      int64           `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ECGData     json.RawMessage `json:"ecg_data"`
}

func (h *TestHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req CreateTestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	test := &models.Test{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		ECGData:     req.ECGData,
		Status:      "pending",
	}

	if err := h.repo.Create(test); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating test")
		return
	}

	respondWithJSON(w, http.StatusCreated, test)
}

func (h *TestHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		respondWithError(w, http.StatusBadRequest, "Test ID is required")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid test ID")
		return
	}

	test, err := h.repo.GetByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Test not found")
		return
	}

	respondWithJSON(w, http.StatusOK, test)
}

func (h *TestHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	tests, err := h.repo.GetByUserID(userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving tests")
		return
	}

	respondWithJSON(w, http.StatusOK, tests)
}

func (h *TestHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var test models.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.repo.Update(&test); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating test")
		return
	}

	respondWithJSON(w, http.StatusOK, test)
}
