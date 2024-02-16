package models

import (
	"errors"

	"github.com/sikehish/Go-Event-Booking-API/db"
	"github.com/sikehish/Go-Event-Booking-API/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error { //Setting user to *User type so that on modifying user.ID, the user entity gets modified and so the response would have the correct ID as opposed to displaying ID:0
	query := `INSERT INTO users(email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPw, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPw)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email=?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		// return err
		return errors.New("Credentials invalid") //If email isnt present
	}

	isValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !isValid {
		return errors.New("Credentials invalid") //If password is incorrect
	}

	return nil

}
