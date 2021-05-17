package routers

import (
	"errors"
	"strings"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email var for endpoint
var Email string

//IdUser var for endpoint
var IdUser string

//function to validate token
func ProcessToken(tk string) (models.Claim, bool, string, error) {
	miKey := []byte("kdñlajidamliwu8daljkdalkjdil")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return *claims, false, string(""), errors.New("Format token invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miKey, nil
	})

	if err == nil {
		_, find, _ := bd.CheckUserExist(claims.Email)
		if find == true {
			Email = claims.Email
			IdUser = claims.ID.Hex()
		}
		return *claims, find, IdUser, nil
	}

	if !tkn.Valid {
		return *claims, false, string(""), errors.New("Invalid token")
	}
	return *claims, false, string(""), err
}
