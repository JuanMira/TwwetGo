package middleware

import (
	"net/http"

	"github.com/JuanMira/tweetgo/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection bd is missing", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
