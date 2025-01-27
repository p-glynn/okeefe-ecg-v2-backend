package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/p-glynn/okeefe-ecg-v2-backend/models"
	"github.com/p-glynn/okeefe-ecg-v2-backend/repository"
)

type CommentHandler struct {
	repo *repository.CommentRepository
}

func NewCommentHandler(db *sql.DB) *CommentHandler {
	return &CommentHandler{
		repo: repository.NewCommentRepository(db),
	}
}

type CreateCommentRequest struct {
	TestID  int64  `json:"test_id" binding:"required"`
	UserID  int64  `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := &models.Comment{
		TestID:  req.TestID,
		UserID:  req.UserID,
		Content: req.Content,
	}

	if err := h.repo.Create(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) GetByTest(c *gin.Context) {
	testID, err := strconv.ParseInt(c.Param("test_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	comments, err := h.repo.GetByTestID(testID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ID = id
	if err := h.repo.Update(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating comment"})
		return
	}

	c.JSON(http.StatusOK, comment)
}
