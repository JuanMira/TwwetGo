package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/jwt"
	"github.com/JuanMira/tweetgo/models"
)

// function to login in the system
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data invalid "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	document, exist := bd.TryLogin(t.Email, t.Password)

	if exist == false {
		http.Error(w, "Data invalid", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "A error has ocurred with token", 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
