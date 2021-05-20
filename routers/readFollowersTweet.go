package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JuanMira/tweetgo/bd"
)

func ReadFollowersTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Param page should be send", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Param page should be a integer", 400)
		return
	}

	response, status := bd.ReadFollowersTweets(IdUser, page)
	if status == false {
		http.Error(w, "Error read tweets", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
