package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JuanMira/tweetgo/bd"
)

func ListUser(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Param page must be major to 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	result, status := bd.ReadUsersAll(IdUser, pag, search, typeUser)

	if status == false {
		http.Error(w, "Something was wrong", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
