package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/JuanMira/tweetgo/bd"
	"github.com/JuanMira/tweetgo/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileUpload string = "uploads/avatars/" + IdUser + "." + extension

	f, err := os.OpenFile(fileUpload, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error uploading the image", 400)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error uploading the image "+err.Error(), 400)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IdUser + "." + extension
	status, err = bd.ModifyProfile(user, IdUser)

	if err != nil || status == false {
		http.Error(w, "Error uploading avatar in the database", 400)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
}
