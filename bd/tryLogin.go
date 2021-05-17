package bd

import (
	"github.com/JuanMira/tweetgo/models"
	"golang.org/x/crypto/bcrypt"
)

//check login data in bd
func TryLogin(email string, password string) (models.User, bool) {
	user, find, _ := CheckUserExist(email)

	if find == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
