package routers

import (
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var t models.Relation

	t.UserId = IdUser
	t.UserRelationId = id

	status, err := bd.DeleteRelation(t)

	if err != nil {
		http.Error(w, "Something was wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Error unfollow the dude "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
