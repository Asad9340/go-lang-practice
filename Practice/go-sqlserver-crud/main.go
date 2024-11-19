package main

import (
	"fmt"
	"go-sqlserver-crud/config"
	"go-sqlserver-crud/models"
	routes "go-sqlserver-crud/routers"
	"net/http"
)

func main() {
	// Connect to the Database
	config.ConnectDB()

	// Auto-migrate the User model
	config.DB.AutoMigrate(&models.User{})

	// Set up the router
	router := routes.RegisterRoutes()

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
