package handler

import (
	"Forum/backend/db"
	"Forum/backend/structs" // Assurez-vous d'importer le package contenant les structures (par exemple, models)
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

// generateSessionToken génère un token de session aléatoire
func generateSessionToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(token)
}

// CreateSession crée une nouvelle session pour l'utilisateur
func CreateSession(userID uint) (string, error) {
	sessionToken := generateSessionToken()
	expiration := time.Now().Add(24 * time.Hour)

	// Créer une nouvelle session
	session := structs.Session{
		UserID:      userID,
		SessionToken: sessionToken,
		ExpiresAt:   expiration,
	}

	// Insérer la session dans la base de données
	result := db.DB.Create(&session)
	if result.Error != nil {
		return "", result.Error
	}

	return sessionToken, nil
}

// Login gère la logique de connexion de l'utilisateur
func Login(w http.ResponseWriter, r *http.Request) {
	// Ici, vous devez vérifier les identifiants de l'utilisateur et récupérer son userID
	// Supposons que vous avez récupéré le userID de l'utilisateur après la validation des identifiants

	var userID uint // Remplacez cette ligne par la récupération du userID réel
	// Exemple: userID = 1

	// Créer une session pour l'utilisateur
	sessionToken, err := CreateSession(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
		return
	}

	// Créer un cookie avec le token de session
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// Rediriger l'utilisateur vers la page d'accueil
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// Register gère l'inscription des utilisateurs
func Register(w http.ResponseWriter, r *http.Request) {
	// Crée un nouveau compte et l'ajoute à la base de données
	// Vous devez ajouter la logique pour récupérer les données du formulaire et enregistrer l'utilisateur dans la base de données
	fmt.Fprintln(w, "Page Register")
}
