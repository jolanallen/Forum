package middlewares


import (
	"net/http"
	"sync"
	"time"
	"log"
)

var mu sync.Mutex
var ipRequests = make(map[string][]time.Time)

// Define rate limit parameters
const maxRequests = 10            // Max requests per IP
const timeWindow = 1 * time.Minute // Time window for rate limiting

// RateLimit middleware to limit requests and redirect if exceeded
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		ip := r.RemoteAddr // Get client IP address
		now := time.Now()

		// Clean up old request timestamps
		ipRequests[ip] = filterOldRequests(ipRequests[ip], now, timeWindow)

		// Check if rate limit exceeded
		if len(ipRequests[ip]) >= maxRequests {
			// Log the excess traffic and redirect the user to an overload page
			log.Printf("Rate limit exceeded for IP: %s", ip)

			// Redirect to a "Too Many Requests" or "Overload" page (or maintenance page)
			http.Redirect(w, r, "/overload", http.StatusTooManyRequests)
			return
		}

		// Log the request time
		ipRequests[ip] = append(ipRequests[ip], now)

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

// Helper function to filter out old requests that are outside the time window
func filterOldRequests(timestamps []time.Time, now time.Time, window time.Duration) []time.Time {
	var validTimestamps []time.Time
	for _, timestamp := range timestamps {
		if now.Sub(timestamp) <= window {
			validTimestamps = append(validTimestamps, timestamp)
		}
	}
	return validTimestamps
}
