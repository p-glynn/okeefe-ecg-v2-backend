package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"server/models"
	"server/repository"
)

type CommentHandler struct {
	repo *repository.CommentRepository
}

func NewCommentHandler(repo *repository.CommentRepository) *CommentHandler {
	return &CommentHandler{repo: repo}
}

type CreateCommentRequest struct {
	TestID  int64  `json:"test_id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment := &models.Comment{
		TestID:  req.TestID,
		UserID:  req.UserID,
		Content: req.Content,
	}

	if err := h.repo.Create(comment); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating comment")
		return
	}

	respondWithJSON(w, http.StatusCreated, comment)
}

func (h *CommentHandler) GetByTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	testIDStr := r.URL.Query().Get("test_id")
	if testIDStr == "" {
		respondWithError(w, http.StatusBadRequest, "Test ID is required")
		return
	}

	testID, err := strconv.ParseInt(testIDStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid test ID")
		return
	}

	comments, err := h.repo.GetByTestID(testID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error retrieving comments")
		return
	}

	respondWithJSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.repo.Update(&comment); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating comment")
		return
	}

	respondWithJSON(w, http.StatusOK, comment)
}
