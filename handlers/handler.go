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
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
