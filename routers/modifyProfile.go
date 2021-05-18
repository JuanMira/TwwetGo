package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

var state bool

//route endopoint to modify profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Wrong data", http.StatusBadRequest)
		return
	}

	state, err = bd.ModifyProfile(t, IdUser)

	if err != nil {
		http.Error(w, "Something was wrong, try it again "+err.Error(), 400)
		return
	}

	if state == false {
		http.Error(w, "Nothing happens with modifiy user"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
