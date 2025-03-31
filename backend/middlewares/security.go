package middlewares

import (
	"Forum/backend/handler"
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var mu sync.Mutex
var loginLimits = make(map[string]*rate.Limiter) // Anti-brute-force pour /auth/login et /auth/register
var globalLimits = make(map[string]*rate.Limiter) // limites global

// Fonction pour récupérer un rate limiter ou en créer un
func getLimiter(ip string, limitMap map[string]*rate.Limiter, rps rate.Limit, burst int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limitMap[ip]
	if !exists {
		limiter = rate.NewLimiter(rps, burst)
		limitMap[ip] = limiter
	}
	return limiter
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		var limiter *rate.Limiter

		switch r.URL.Path {
		case "/auth/login", "/auth/register":
			limiter = getLimiter(ip, loginLimits, 3, 2) // 3 requêtes/sec, capacité de 2 requêtes du tampon (burst)
		default:
			limiter = getLimiter(ip, globalLimits, 20, 5) // 20 requêtes/sec, tampon de 5
		}

		// Vérification du rate limit
		if !limiter.Allow() {
			log.Printf("[!!!!!! ALERT !!!!] Request blocked from IP: %s | URL: %s | Method: %s | UserAgent: %s \n", ip, r.URL.Path, r.Method, r.UserAgent())
				handler.OverloadHandler(w, r)
			return
		}

		next.ServeHTTP(w, r) 
	})
}
