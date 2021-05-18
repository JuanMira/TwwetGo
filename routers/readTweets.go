package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JuanMira/tweetgo/bd"
)

func RetrieveTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Param id is required", 400)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Param page is required", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Param must be greater than 0", 400)
		return
	}

	pag := int64(page)

	response, correct := bd.RetrieveTweets(ID, pag)

	if correct == false {
		http.Error(w, "Error to retrieve the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
