package routes

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"okeefe-ecg-v2-backend/repository"
)

var router = gin.Default()
var userRoutes *UserRoutes

func Run() {
	// Initialize database connection
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userRoutes = NewUserRoutes(userRepo)
	getRoutes()
	router.Run()
}

func getRoutes() {
	root := &router.RouterGroup
	userRoutes.addUserRoutes(root)
	addTestsRoutes(root)
}