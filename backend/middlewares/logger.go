package middlewares

import (
	"net/http"
	"log"
	"strings"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		ip := r.RemoteAddr
		if strings.Contains(r.Header.Get("X-Forwarded-For"), ",") {
			ip = strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
		}
		log.Printf("[Acess] => IP: %s | Method: %s | URL: %s | User-Agent: %s", ip, r.Method, r.URL.Path, r.UserAgent())
		next.ServeHTTP(w, r)
=======
		///////////////////////////logic de création de log pour chaque requetes effectuer //////////////////////////////////////////
		next.ServeHTTP(w, r) // Passer au handler suivant //
>>>>>>> c2ac08b (first commit features)
	})
}
