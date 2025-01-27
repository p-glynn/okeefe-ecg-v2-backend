package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/p-glynn/okeefe-ecg-v2-backend/repository"
)

var router = gin.Default()
var userRoutes *UserRoutes

func Run() {
	userRepo := repository.NewUserRepository()
	userRoutes = NewUserRoutes(userRepo)
	getRoutes()
	router.Run()
}

func getRoutes() {
	root := &router.RouterGroup
	userRoutes.addUserRoutes(root)
	addTestsRoutes(root)
}