package jwt

import (
	"time"

	"github.com/JuanMira/tweetgo/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//create a token
func GenerateJWT(t models.User) (string, error) {
	miClave := []byte("kdñlajidamliwu8daljkdalkjdil")
	payload := jwt.MapClaims{
		"Email":     t.Email,
		"Name":      t.Name,
		"LastName":  t.LastName,
		"BirthDate": t.BirthDate,
		"Location":  t.Location,
		"WebSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	//formato

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
