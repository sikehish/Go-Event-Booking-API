package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/models"
	"github.com/sikehish/Go-Event-Booking-API/utils"
)

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

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event) //event is populated with data. It doesnt return any error if any field is missing, but using struct tags we can enforce ShouldBindJSON to return an error if a field is missing

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data parsing failed :("})
		return
	}

	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event, Try again later :("})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing event id failed"})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event doesnt exist",
		})
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data parsing failed :("})
		return
	}

	updatedEvent.ID = eventId //eventId is a param while updatedEvent is obtained the body of the request

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event, Try again later :("})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing event id failed"})
		return
	}

	deletedEvent, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event doesnt exist",
		})
	}

	err = deletedEvent.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event, Try again later :("})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted", "deletedEvent": deletedEvent})

}
