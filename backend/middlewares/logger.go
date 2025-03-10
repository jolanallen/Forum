package middlewares

import (
	"net/http"
)



func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		///////////////////////////logic de création de log pour chaque requetes effectuer //////////////////////////////////////////
		next.ServeHTTP(w, r) // Passer au handler suivant //
	})
}