package middlewares

import (
	"net/http"
	"net"
	"sync"


	"golang.org/x/time/rate"
)

var visitors = make (map[string]*rate.Limiter)
var mu sync.Mutex

// getVisitorLimiter retrieves the rate limiter for a given IP (or creates one if needed)
func getVisitorLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		// 20 requests per second with a burst of 5
		limiter = rate.NewLimiter(20, 5)
		visitors[ip] = limiter
	}
	return limiter
}



// RateLimit is the middleware function
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err !=nil {
			http.Error(w, "Unable to determine IP address", http.StatusInternalServerError)
			return
		}

		limiter := getVisitorLimiter(ip)
		if !limiter .Allow() {
			http.Error(w," Too many requests", http.StatusTooManyRequests)
			return
		} 
		next.ServeHTTP(w, r) // Passer au handler suivant
	})
}