package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sikehish/Go-Event-Booking-API/db"
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
func VerifyToken(token string) (int64, error) {
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
		return 0, errors.New("Could not parse token")
	}

	isValid := parsedToken.Valid

	if !isValid {
		return 0, errors.New("Invalid token!")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) //claims will be of type jwt.MapClaims if ok=true

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)
	userId, ok := claims["userId"].(int64)
	//checking for string and int64 is optional as we know for sure that weve set both the fields to string and int64 types respectively

	if !ok {
		return 0, errors.New("userId is invalid")
	}

	//Checking the database to see if the userId exists(Sometimes token may be valid but userId may not be there in the Table)

	query := `SELECT id FROM users WHERE id= ?`

	row := db.DB.QueryRow(query, userId)

	var id int64
	err = row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
