package utils

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var Templates *template.Template

var F = &structs.Forum{} ///variables global ( à voir ce que je voulais en faire)

// création d'un token de sessions aléatoire de byte en hexa
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
// le pb c'est qu'il prend par rapport à la session de son nav et pas de notre bdd
func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0 // Pas de session
	}
	var userID uint64
	fmt.Sscanf(cookie.Value, "%d", &userID)
	return userID
}
// je pourrais l'enlevé si JOLAN VIENS M4AIDER
func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du parsing des templates:", err)
	}
}

// vérification du hash = hash du mot de passe 
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
//pour hashé le mdp
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}



// pour les COOKIES
func CheckSession(next http.HandlerFunc) http.HandlerFunc {
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
func CheckAdmin(next http.HandlerFunc) http.HandlerFunc {
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

// CreateSession crée une nouvelle session pour l'utilisateur
func CreateSession(userID uint64) (string, error) {
	sessionToken := GenerateToken()
	expiration := time.Now().Add(24 * time.Hour)

	// Créer une nouvelle session
	session := structs.Session{
		UserID:       userID,
		SessionToken: sessionToken,
		ExpiresAt:    expiration,
	}

	// Insérer la session dans la base de données
	result := db.DB.Create(&session)
	if result.Error != nil {
		return "", result.Error
	}

	return sessionToken, nil
}














// Récupérer tous les users mais rien avec le front
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(users)
}

// Récupérer tous les admins
// ça ne return rien, c'est juste pour la prog
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
}
