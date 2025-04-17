package middlewares

import (
	"log"
	"net"
	"net/http"
	"sync"
	"golang.org/x/time/rate"
)

var mu sync.Mutex
// A map to store rate limiters for login-related routes (anti-brute-force for login and register)
var loginLimits = make(map[string]*rate.Limiter) 
// A map to store rate limiters for global routes (used for all other requests)
var globalLimits = make(map[string]*rate.Limiter) 

// getLimiter retrieves or creates a rate limiter for a given IP address.
// It uses a provided limit map (loginLimits or globalLimits), requests per second (rps), and burst capacity.
func getLimiter(ip string, limitMap map[string]*rate.Limiter, rps rate.Limit, burst int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	// Check if a limiter already exists for the provided IP address.
	limiter, exists := limitMap[ip]
	if !exists {
		// If it doesn't exist, create a new limiter with the provided rate limit (rps) and burst size.
		limiter = rate.NewLimiter(rps, burst)
		limitMap[ip] = limiter // Store the limiter in the map for future use.
	}
	return limiter // Return the rate limiter.
}

// RateLimit is a middleware that enforces rate limiting based on the IP address of the client.
// It applies different limits for login-related routes and global routes.
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the IP address from the request's remote address.
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		// Declare a variable to hold the rate limiter based on the request path.
		var limiter *rate.Limiter

		// Apply different rate limits based on the requested URL path.
		switch r.URL.Path {
		// For login and register related routes, use stricter limits to prevent brute-force attacks.
		case "/auth/login", "/auth/register", "/login", "/register":
			limiter = getLimiter(ip, loginLimits, 555, 2555) // 5 requests/sec, burst capacity of 2555 requests
		// For all other routes, apply a more lenient global limit.
		default:
			limiter = getLimiter(ip, globalLimits, 205, 5555) // 20 requests/sec, burst capacity of 5555 requests
		}

		// Check if the rate limiter allows the request (i.e., within the rate limit).
		if !limiter.Allow() {
			// If the rate limit is exceeded, log the event and return a 429 Too Many Requests error.
			log.Printf("[!!!!!! ALERT !!!!] Request blocked from IP: %s | URL: %s | Method: %s | UserAgent: %s", ip, r.URL.Path, r.Method, r.UserAgent())
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return // Stop further handling of the request.
		}

		// If the request is allowed, pass it to the next handler in the middleware chain.
		next.ServeHTTP(w, r)
	})
}
