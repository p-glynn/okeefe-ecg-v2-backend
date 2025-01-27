package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/p-glynn/okeefe-ecg-v2-backend/models"
	"github.com/p-glynn/okeefe-ecg-v2-backend/repository"
)

type TestHandler struct {
	repo *repository.TestRepository
}

func NewTestHandler(db *sql.DB) *TestHandler {
	return &TestHandler{
		repo: repository.NewTestRepository(db),
	}
}

type CreateTestRequest struct {
	UserID      int64           `json:"user_id" binding:"required"`
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description"`
	ECGData     json.RawMessage `json:"ecg_data" binding:"required"`
}

func (h *TestHandler) Create(c *gin.Context) {
	var req CreateTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating test"})
		return
	}

	c.JSON(http.StatusCreated, test)
}

func (h *TestHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	test, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Test not found"})
		return
	}

	c.JSON(http.StatusOK, test)
}

func (h *TestHandler) GetByUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	tests, err := h.repo.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tests"})
		return
	}

	c.JSON(http.StatusOK, tests)
}

func (h *TestHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	test.ID = id
	if err := h.repo.Update(&test); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating test"})
		return
	}

	c.JSON(http.StatusOK, test)
}
