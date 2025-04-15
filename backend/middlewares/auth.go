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
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        // Vérification de la session utilisateur
        var userSession structs.SessionUser
        result := db.DB.Where("sessionToken = ?", cookie.Value).First(&userSession)
        if result.Error != nil || userSession.ExpiresAt.Before(time.Now()) {
            // Vérification de la session admin
            var adminSession structs.SessionAdmin
            result = db.DB.Where("sessionToken = ?", cookie.Value).First(&adminSession)
            if result.Error != nil || adminSession.ExpiresAt.Before(time.Now()) {
                // Vérification de la session invité
                var guestSession structs.SessionGuest
                result = db.DB.Where("sessionToken = ?", cookie.Value).First(&guestSession)
                if result.Error != nil || guestSession.ExpiresAt.Before(time.Now()) {
                    http.Redirect(w, r, "/login", http.StatusSeeOther)
                    return
                }

                // Invité : Stockage dans le contexte
                ctx := context.WithValue(r.Context(), "guestID", guestSession.GuestID)
                next.ServeHTTP(w, r.WithContext(ctx))
                return
            }

            // Admin : Stockage dans le contexte
            ctx := context.WithValue(r.Context(), "adminID", adminSession.AdminID)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }

        // Utilisateur : Stockage dans le contexte
        ctx := context.WithValue(r.Context(), "userID", userSession.UserID)
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
