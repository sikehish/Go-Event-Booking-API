package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //Or you can have an anonymous function in place ofgetEvents
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
}
