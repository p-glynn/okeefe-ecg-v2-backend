package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/p-glynn/okeefe-ecg-v2-backend/db"
	"github.com/p-glynn/okeefe-ecg-v2-backend/handlers"
)

func main() {
	fmt.Println("Starting server...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default or system environment")
	}

	// Initialize database connection
	dbConfig := db.NewConfig()
	database, err := dbConfig.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Set Gin to release mode in production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))

	// Initialize API routes
	api := router.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			userHandler := handlers.NewUserHandler(database)
			users.POST("", userHandler.Create)
			users.GET("/:id", userHandler.Get)
			users.PUT("/:id", userHandler.Update)
		}

		// Test routes
		tests := api.Group("/tests")
		{
			testHandler := handlers.NewTestHandler(database)
			tests.POST("", testHandler.Create)
			tests.GET("/:id", testHandler.Get)
			tests.GET("/user/:user_id", testHandler.GetByUser)
			tests.PUT("/:id", testHandler.Update)
		}

		// Comment routes
		comments := api.Group("/comments")
		{
			commentHandler := handlers.NewCommentHandler(database)
			comments.POST("", commentHandler.Create)
			comments.GET("/test/:test_id", commentHandler.GetByTest)
			comments.PUT("/:id", commentHandler.Update)
		}
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}