package middlewares

import (
	"net/http"
)

///////////////////////rates limiting /////////////////////////
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		////////////////////////// logic  du rate limiting ////////////////////////////////////////////////////
		
		next.ServeHTTP(w, r) // Passer au handler suivant
	})
}