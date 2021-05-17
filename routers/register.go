package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

/* function to register a user in bd */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.User
	//solo se puede usar un body
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in data "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be than 6 characters", 400)
		return
	}

	_, find, _ := bd.CheckUserExist(t.Email)

	if find == true {
		http.Error(w, "Email already exist", 400)
		return
	}

	_, status, bdErr := bd.InsertData(t)
	if bdErr != nil {
		http.Error(w, "A error has ocurred trying create new user ", 400)
		return
	}

	if status == false {
		http.Error(w, "Not is posible create a new user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
