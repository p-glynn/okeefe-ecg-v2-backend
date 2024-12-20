package server

import (
	"fmt"
	"routes"
)

func main() {
	fmt.Println("Starting server...")
	routes.Run()
}