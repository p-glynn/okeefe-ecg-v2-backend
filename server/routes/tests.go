package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTestsRoutes(rg *gin.RouterGroup) {
	tests := rg.Group("/tests")

	tests.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "tests")
	})
	tests.GET("/comments", func(c *gin.Context) {
		c.JSON(http.StatusOK, "tests comments")
	})
	tests.GET("/pictures", func(c *gin.Context) {
		c.JSON(http.StatusOK, "tests pictures")
	})
}