package utils

import "golang.org/x/crypto/bcrypt"

//Bcrypt uses a cost parameter that specify the number of cycles to use in the algorithm. Increasing this number the algorithm will spend more time to generate the hash output. The cost parameter is represented by an integer value between 4 to 31.A cost is a measure of how many times to run the hash -- how slow it is. You want it to be slow. Again, this is a redundant layer of security for if the hashed passwords are stolen. It makes it prohibitively expensive to brute-force anything.

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
