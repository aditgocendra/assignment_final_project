package main

import (
	"final_project/database"
	"final_project/router"
)

// @title       Final Project API
// @version     1.0
// @description This is a assignment for DTS-Hacktiv8 final project.

// @contact.name  Aditya Gocendra
// @contact.email gocendra123@gmail.com

// @host localhost:8080

// @securityDefinitions.apiKey Bearer <JWT>
// @in                         header
// @name                       Authorization

func main() {
	database.StartDB()

	r := router.StartApp()
	r.Run(":8080")
}