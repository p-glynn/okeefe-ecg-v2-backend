package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"okeefe-ecg-v2-backend/repository"
)

type UserRoutes struct {
	userRepo *repository.UserRepository
}

func NewUserRoutes(userRepo *repository.UserRepository) *UserRoutes {
	return &UserRoutes{userRepo: userRepo}
}

func (ur *UserRoutes) addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users comments")
	})
	users.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users pictures")
	})

	users.GET("/:id", func(c *gin.Context) {
		userID := c.Param("id")
		id, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}
		
		user, err := ur.userRepo.GetByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
