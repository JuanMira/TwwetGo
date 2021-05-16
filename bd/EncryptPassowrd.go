package bd

import "golang.org/x/crypto/bcrypt"

// function to encrypt password
func EncryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
