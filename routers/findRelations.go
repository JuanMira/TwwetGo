package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func FindRelations(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var t models.Relation

	t.UserId = IdUser
	t.UserRelationId = id

	var resp models.ResponseRelation

	status, err := bd.FindRelations(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
