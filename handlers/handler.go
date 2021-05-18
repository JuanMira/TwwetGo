package handler

import (
	"log"
	"net/http"
	"os"

	mw "github.com/JuanMira/tweetgo/middlewares"
	"github.com/JuanMira/tweetgo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//mw means middleware

//Handler port and server listening
func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", mw.CheckBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", mw.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", mw.CheckBD(mw.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", mw.CheckBD(mw.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	//tweet route
	router.HandleFunc("/insertTweet", mw.CheckBD(mw.ValidateJWT(routers.Tweet))).Methods("POST")
	router.HandleFunc("/readTweets", mw.CheckBD(mw.ValidateJWT(routers.RetrieveTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", mw.CheckBD(mw.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	//avatar and banner
	router.HandleFunc("/uploadAvatar", mw.CheckBD(mw.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", mw.CheckBD(mw.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", mw.CheckBD(mw.ValidateJWT(routers.UploadBanner))).Methods("POSt")
	router.HandleFunc("/getBanner", mw.CheckBD(mw.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/insertRelation", mw.CheckBD(mw.ValidateJWT(routers.InsertRelation))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
