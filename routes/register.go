package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/models"
)

func eventRegister(context *gin.Context) { //User registering for an event
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing event id failed"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user :("})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered for the event successfully"})
}

func eventUnregister(context *gin.Context) { //User unregistering for an event
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Parsing event id failed"})
		return
	}

	// event, err := models.GetEventByID(eventId)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	// 	return
	// }
	// err = event.Register(userId)
	// if err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister user :("})
	// 	return
	// }
	//OR
	var event models.Event
	event.ID = eventId
	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister user :("})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User unregistered for the event successfully"})
}
