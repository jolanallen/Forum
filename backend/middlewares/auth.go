package middlewares

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"database/sql"
	"net/http"
	"time"
)

// Authentication middleware checks if the user is authenticated
// It verifies the session token in the cookie and stores the session details in the context.
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the session token exists in the cookie
		cookie, err := r.Cookie("sessionToken")
		if err != nil || cookie.Value == "" {
			// Redirect to login if the session token is missing or invalid
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Check user session validity
		var userSession structs.SessionUser
		query := "SELECT userID, expiresAt FROM sessionsUsers WHERE sessionToken = ?"
		err = db.DB.QueryRow(query, cookie.Value).Scan(&userSession.UUserID, &userSession.UExpiresAt)
		if err != nil || userSession.UExpiresAt.Before(time.Now()) {
			// Check admin session validity
			var adminSession structs.SessionAdmin
			query := "SELECT adminID, expiresAt FROM sessionsAdmins WHERE sessionToken = ?"
			err = db.DB.QueryRow(query, cookie.Value).Scan(&adminSession.AAdminID, &adminSession.AExpiresAt)
			if err != nil || adminSession.AExpiresAt.Before(time.Now()) {
				// Check guest session validity
				var guestSession structs.SessionGuest
				query := "SELECT guestID, expiresAt FROM sessionsGuests WHERE sessionToken = ?"
				err = db.DB.QueryRow(query, cookie.Value).Scan(&guestSession.GGuestID, &guestSession.GExpiresAt)
				if err != nil || guestSession.GExpiresAt.Before(time.Now()) {
					// Redirect to login if all sessions are invalid or expired
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}

				// Guest: Store guestID in context and pass the request to the next handler
				ctx := context.WithValue(r.Context(), "guestID", guestSession.GGuestID)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// Admin: Store adminID in context and pass the request to the next handler
			ctx := context.WithValue(r.Context(), "adminID", adminSession.AAdminID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// User: Store userID in context and pass the request to the next handler
		ctx := context.WithValue(r.Context(), "userID", userSession.UUserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminAuthorization middleware checks if the authenticated user is an admin
// It ensures that only users with an admin role can access certain routes.
func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get userID from the context, which was set by the Authentication middleware
		userID := r.Context().Value("userID")
		var adminID uint64
		// Query to check if the user is an admin
		query := "SELECT adminID FROM admins WHERE userID = ?"
		err := db.DB.QueryRow(query, userID).Scan(&adminID)
		if err != nil && err != sql.ErrNoRows {
			// Return forbidden error if there's an issue with the query
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// If the user is not an admin (no result or invalid adminID), return forbidden
		if adminID == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		// Pass the request to the next handler if the user is an admin
		next.ServeHTTP(w, r)
	})
}
