package routers

import (
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Param id is missing", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserId = IdUser
	t.UserRelationId = ID

	status, err := bd.InsertRelation(t)

	if err != nil {
		http.Error(w, "Something was wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error follow the dude "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
