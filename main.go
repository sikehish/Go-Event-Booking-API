package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/db"
	"github.com/sikehish/Go-Event-Booking-API/routes"
)

func main() {
	db.InitDB()

	server := gin.Default() // Create a new Gin router

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
