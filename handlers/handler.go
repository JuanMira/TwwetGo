package handler

import (
	"log"
	"net/http"
	"os"

	middleware "github.com/JuanMira/tweetgo/middlewares"
	"github.com/JuanMira/tweetgo/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handler port and server listening
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
