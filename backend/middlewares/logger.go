package middlewares

import (
	"net/http"
	"log"
	"time"
)

///////////////////////////logic de création de log pour chaque requetes effectuer //////////////////////////////////////////

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the duration of the request
		log.Printf("Duration: %v", time.Since(start))
	})
}