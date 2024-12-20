package main

import (
	"fmt"
	"server/routes"
)

func main() {
	fmt.Println("Starting server...")
	routes.Run()
}