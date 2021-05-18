package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/JuanMira/tweetgo/bd"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(w, "Should send param id", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(Id)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error copy image", http.StatusBadRequest)
	}
}
