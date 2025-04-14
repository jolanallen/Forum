package middlewares

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"net/http"
	"time"
)


func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionToken")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
			return
		}

		var session structs.Session
		result := db.DB.Where("sessionToken = ?", cookie.Value).First(&session)
		if result.Error != nil || session.ExpiresAt.Before(time.Now()) {
			http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID")
		var admin structs.Admin
		if err := db.DB.Where("userID = ?", userID).First(&admin).Error; err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
