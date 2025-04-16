package middlewares

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"database/sql"
	"net/http"
	"time"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionToken")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Vérification de la session utilisateur
		var userSession structs.SessionUser
		query := "SELECT userID, expiresAt FROM session_users WHERE sessionToken = ?"
		err = db.DB.QueryRow(query, cookie.Value).Scan(&userSession.UUserID, &userSession.UExpiresAt)
		if err != nil || userSession.UExpiresAt.Before(time.Now()) {
			// Vérification de la session admin
			var adminSession structs.SessionAdmin
			query := "SELECT adminID, expiresAt FROM session_admins WHERE sessionToken = ?"
			err = db.DB.QueryRow(query, cookie.Value).Scan(&adminSession.AAdminID, &adminSession.AExpiresAt)
			if err != nil || adminSession.AExpiresAt.Before(time.Now()) {
				// Vérification de la session invité
				var guestSession structs.SessionGuest
				query := "SELECT guestID, expiresAt FROM session_guests WHERE sessionToken = ?"
				err = db.DB.QueryRow(query, cookie.Value).Scan(&guestSession.GGuestID, &guestSession.GExpiresAt)
				if err != nil || guestSession.GExpiresAt.Before(time.Now()) {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}

				// Invité : Stockage dans le contexte
				ctx := context.WithValue(r.Context(), "guestID", guestSession.GGuestID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Admin : Stockage dans le contexte
			ctx := context.WithValue(r.Context(), "adminID", adminSession.AAdminID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// Utilisateur : Stockage dans le contexte
		ctx := context.WithValue(r.Context(), "userID", userSession.UUserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID")
		var adminID uint64
		query := "SELECT adminID FROM admins WHERE userID = ?"
		err := db.DB.QueryRow(query, userID).Scan(&adminID)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Si l'utilisateur n'est pas un admin (aucun résultat ou erreur)
		if adminID == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
