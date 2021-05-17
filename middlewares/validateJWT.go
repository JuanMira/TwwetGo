package middleware

import (
	"net/http"

	"github.com/JuanMira/tweetgo/routers"
)

//validation jwt
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Token wrong ", http.StatusBadRequest)
		}
		next.ServeHTTP(rw, r)
	}
}
