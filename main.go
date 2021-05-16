package main

import (
	"log"

	"github.com/JuanMira/tweetgo/bd"
	handler "github.com/JuanMira/tweetgo/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("BD not connected")
		return
	}

	handler.Handlers()
}
