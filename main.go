package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/models"
)

func main() {
	server := gin.Default()          // Create a new Gin router
	server.GET("/events", getEvents) //Or you can have an anonymous function in place ofgetEvents
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{ //http.StatusOK is the same as Status code 200. You can use it interchangably
	// 	"message": "Hello!",
	// })

	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) //event is populated with data. It doesnt return any error if any field is missing, but sing struct tags we can enforce ShouldBindJSON to return an error if a field is missing

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
