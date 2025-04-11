package middlewares

import (
	"net/http"
)

// Middleware d'authentification
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier si l'utilisateur est connecté (exemple simple avec cookie/session)
		///A COMPLETER AVEC LE CHECK DE SESSION DE COOKIES SI L'UTILISATERU EST BIEN UN UTILISATEUR CONNECTER//////////////////////////////: ///

////////////////////vas voir dans le validationService ///////////////////////////////////
//checksessions surement ( nan mais aussi authentification de quoi, si tu précise pas je fais du mieux que je peux)

		next.ServeHTTP(w, r) // Passer au handler suivant
	})
}


// Middleware d'autorisation admin
func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier si l'utilisateur est admin (exemple simple)
		////////:A COMPLETER IL FAUT VERIFIER AVEC UN COOOKIES QUE LA PERSOONNES CONNECTER EST BIEN ADMIN/////////////////////////////

////////////////////vas voir dans le validationService ///////////////////////////////////
//j'ai surement mis checkadmin

		next.ServeHTTP(w, r)
	})
}