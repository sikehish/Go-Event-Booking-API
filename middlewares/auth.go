package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sikehish/Go-Event-Booking-API/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
		// AbortWithStatusJSON calls `Abort()` and then `JSON` internally. This method stops the chain, writes the status code and return a JSON body.

		//We dont use context.JSON() as other request handlers in line will still be executed even after sending a response in in this middleware, and we want to ensure that the execution is completely stopped as the user isnt authorized, and that's why we use AbortWithStatusJSON
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}

	context.Set("userId", userId) //Attaching a value to the context which can be accessed whereever context is available
	context.Next()                //Ensures that the next request handler in line would be executed
}
