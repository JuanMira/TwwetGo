package routers

import (
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Param id is missing", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, IdUser)
	if err != nil {
		http.Error(w, "Something was wrong "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
