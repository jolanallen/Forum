package utils

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"net/http"
	"time"
)

// pour les COOKIES
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.Cookie c'est les informations directe récuperer par le navigateur
		cookie, err := r.Cookie("sessionToken")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		var session structs.Session
		result := db.DB.Where("sessionToken = ?", cookie.Value).First(&session)
		if result.Error != nil || session.ExpiresAt.Before(time.Now()) {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// POUR LES ADMIN
// on verifie que l'userID est ou n'est pas un admin
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID")
		var admin structs.Admin
		if err := db.DB.Where("userID = ?", userID).First(&admin).Error; err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
