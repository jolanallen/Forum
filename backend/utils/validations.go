package utils

import (
	"fmt"
	"net/http"
	"time"
)

// Fonction pour initialiser une session (par exemple, lors de la connexion)
func SetSession(w http.ResponseWriter, userID uint64) {
	// Définir un cookie de session
	expiration := time.Now().Add(24 * time.Hour) // Le cookie expire dans 24h
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   string(userID),
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)
}

// Fonction pour récupérer l'ID de l'utilisateur depuis le cookie
func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return 0 // Pas de session
	}

	// Convertir la valeur du cookie en uint64 (c'est un exemple, tu peux adapter selon ton type de session)
	var userID uint64
	fmt.Sscanf(cookie.Value, "%d", &userID)
	return userID
}
