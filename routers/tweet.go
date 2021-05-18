package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func Tweet(w http.ResponseWriter, r *http.Request) {
	var message models.TweetModel
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.Tweet{
		UserId:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)

	if err != nil {
		http.Error(w, "Something was wrong try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Imposible to insert the tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
