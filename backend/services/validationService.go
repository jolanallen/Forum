package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

func checkIfEmailExists(email string) error {
	existingUser, err := GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("Un compte avec cet email existe déjà")
	}
	return nil
}

func GetUserByEmail(email string) (*structs.User, error) {
	var user structs.User
	result := db.DB.Where("userEmail = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func validateRegistrationForm(username, email, password, confirmPassword string) error {
	if username == "" || email == "" || password == "" || confirmPassword == "" {
		return fmt.Errorf("Tous les champs doivent être remplis")
	}

	if password != confirmPassword {
		return fmt.Errorf("Les mots de passe ne correspondent pas")
	}

	return nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// pour hashé le mdp
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// récupération formulaire
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")

		var user structs.User
		//vérifier pour le .ERROR
		if err := db.DB.Where("userEmail = ?", email).First(&user).Error; err != nil {
			//erreur 401
			http.Error(w, "Utilisateur inconnu", http.StatusUnauthorized)
			return
		}

		if !CheckPasswordHash(password, user.UserPasswordHash) {
			//erreur 401 si pwd par bon
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}
		//crée un token de session pour l'utilisateur
		sessionToken, err := CreateSession(user.UserID)
		if err != nil {
			http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
			return
		}

		//on l'insert dans le nav du client (cookie)
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionToken",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})
		//redirection vers home et http.StatusSeeOther sert au cas où il y aura rafraichissement de la page ( status de réussite code 303)
		http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
	} else {
		Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", nil)
	}
}
