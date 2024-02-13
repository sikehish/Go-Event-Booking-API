package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // Create a new Gin router
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{ //http.StatusOK is the same as Status code 200. You can use it interchangably
		"message": "Hello!",
	})
}
