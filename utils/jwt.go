package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "abc123" //temporary

func GenToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))

}

// Adding token verification
func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		//Type Checking
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //if ok is true, then token is indeed signed with HMAC
		//Also jwt.SigningMethodHS256 is of *jwt.SigningMethodHMAC

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return errors.New("Could not parse token")
	}

	isValid := parsedToken.Valid

	if !isValid {
		return errors.New("Invalid token!")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims) //claims will be of type jwt.MapClaims if ok=true

	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	// //checking for string and int64 is optional as we know for sure that weve set both the fields to string and int64 types respectively
	return nil
}
