package middlewares

import (
	"net/http"
	"log"
	"strings"
)

// Logger middleware logs details about incoming requests
// It records the client IP address, HTTP method, request URL, and user-agent.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the client's IP address from the "RemoteAddr" or "X-Forwarded-For" header
		ip := r.RemoteAddr
		if strings.Contains(r.Header.Get("X-Forwarded-For"), ",") {
			// If multiple IPs are listed in the X-Forwarded-For header, get the first one
			ip = strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
		}

		// Log the request details: IP, HTTP method, URL path, and User-Agent
		log.Printf("[Access] => IP: %s | Method: %s | URL: %s | User-Agent: %s", ip, r.Method, r.URL.Path, r.UserAgent())

		// Pass the request to the next handler in the middleware chain
		next.ServeHTTP(w, r)
	})
}
