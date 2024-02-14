package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/models"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data parsing failed :("})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user, Try again later :("})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}
