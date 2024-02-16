package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //Or you can have an anonymous function in place ofgetEvents
	server.GET("/events/:id", getEvent)

	//Registering a middleware
	authGroup := server.Group("/") //creates a new router group. You should add all the routes that have common middlewares or the same path prefix

	authGroup.Use(middlewares.Authenticate)

	authGroup.PUT("/events/:id", updateEvent)
	authGroup.DELETE("/events/:id", deleteEvent)
	authGroup.POST("/events", createEvent)

	authGroup.POST("/signup", signUp)
	authGroup.POST("/login", logIn)

	// //Another way of registering a middleware
	// server.POST("/events", middlewares.Authenticate, createEvent)
}
