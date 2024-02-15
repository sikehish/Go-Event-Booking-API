package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "abc123" //temporary

func GenToken(email string, userId int64) (secret, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(secret)

}
