package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/db"
	"github.com/sikehish/Go-Event-Booking-API/models"
)

func main() {
	db.InitDB()

	server := gin.Default() // Create a new Gin router

	server.GET("/events", getEvents) //Or you can have an anonymous function in place ofgetEvents
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	// context.JSON(http.StatusOK, gin.H{ //http.StatusOK is the same as Status code 200. You can use it interchangably
	// 	"message": "Hello!",
	// })

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Try again later :("})
		return
	}

	if len(events) == 0 {
		context.JSON(http.StatusOK, gin.H{"message": "No events found. Create an event!"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing event id failed"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event doesnt exist",
		})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event) //event is populated with data. It doesnt return any error if any field is missing, but using struct tags we can enforce ShouldBindJSON to return an error if a field is missing

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data parsing failed :("})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event, Try again later :("})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}
