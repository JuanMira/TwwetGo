package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
)

//Show profile data
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "id should be send", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(id)

	if err != nil {
		http.Error(w, "An error ocurred trying find the profile "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
