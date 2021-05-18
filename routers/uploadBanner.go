package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func uploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileUpload string = "uploads/banners/" + IdUser + "." + extension

	f, err := os.OpenFile(fileUpload, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error uploading the banner", 400)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error uploading the banner "+err.Error(), 400)
		return
	}

	var user models.User
	var status bool

	user.Banner = IdUser + "." + extension
	status, err = bd.ModifyProfile(user, IdUser)

	if err != nil || status == false {
		http.Error(w, "Error uploading banner in the database", 400)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
}
