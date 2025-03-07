package middlewares

import (
	"net/http"
)

// Middleware d'authentification
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier si l'utilisateur est connecté (exemple simple avec cookie/session)
		///A COMPLETER AVEC LE CHECK DE SESSION DE COOKIES SI L'UTILISATERU EST BIEN UN UTILISATEUR CONNECTER//////////////////////////////: ///
		next.ServeHTTP(w, r) // Passer au handler suivant
	})
}


// Middleware d'autorisation admin
func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier si l'utilisateur est admin (exemple simple)
		////////:A COMPLETER IL FAUT VERIFIER AVEC UN COOOKIES QUE LA PERSOONNES CONNECTER EST BIEN ADMIN/////////////////////////////
		next.ServeHTTP(w, r)
	})
}