package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)
// /services
// func GoogleRegister; Login
// crée une session pour un utilisateur donné et génère un token de session
func CreateSession(userID uint64) (string, error) {
	sessionToken := GenerateToken()
	expiration := time.Now().Add(24 * time.Hour)

	session := structs.Session{
		UserID:       userID,
		SessionToken: sessionToken,
		ExpiresAt:    expiration,
	}

	result := db.DB.Create(&session)
	if result.Error != nil {
		return "", result.Error
	}

	return sessionToken, nil
}
// /services
// func CreateSession
// crée un token de 32 octets aléatoires pour la session de l'utilisateur
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
// backend\middlewares\auth.go avec le nom Authentication
// verifie si l'utilisateur est authentifié en vérifiant la présence d'un cookie de session valide
func CheckSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.Cookie c'est les informations directe récuperer par le navigateur
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
	}
}
