package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {
	getRoutes()
	router.Run()
}

func getRoutes() {
	root := &router.RouterGroup
	addUserRoutes(root)
	addTestsRoutes(root)
}