package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"           // <- pour oauth2.Config et AccessTypeOffline
	"golang.org/x/oauth2/google"    // <- pour google.Endpoint
	oauth2api "google.golang.org/api/oauth2/v2"
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

func CheckIfEmailExists(email string) error {
	existingUser, err := GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("Un compte avec cet email existe déjà")
	}
	return nil
}

func CheckRegistrationForm(username, email, password, confirmPassword string) error {
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")

		var user structs.User
		if err := db.DB.Where("userEmail = ?", email).First(&user).Error; err != nil {
			http.Error(w, "Utilisateur inconnu", http.StatusUnauthorized)
			return
		}

		if !CheckPasswordHash(password, user.UserPasswordHash) {
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}

		sessionToken, err := CreateSession(user.UserID)
		if err != nil {
			http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "sessionToken",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
	} else {
		Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", nil)
	}
}

var oauth2Config = oauth2.Config{
	ClientID:     "YOUR_GOOGLE_CLIENT_ID",
	ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code manquant", http.StatusBadRequest)
		return
	}

	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'échange du code : %v", err), http.StatusInternalServerError)
		return
	}

	client := oauth2Config.Client(r.Context(), token)

	oauth2Service, err := oauth2api.New(client)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la création du service OAuth2 : %v", err), http.StatusInternalServerError)
		return
	}


	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des infos utilisateur : %v", err), http.StatusInternalServerError)
		return
	}

	var user structs.User
	if err := db.DB.Where("userEmail = ?", userInfo.Email).First(&user).Error; err != nil {
		// Si l'utilisateur n'existe pas, créez un nouvel utilisateur
		newUser := structs.User{
			UserUsername:     userInfo.Name,
			UserEmail:        userInfo.Email,
			UserPasswordHash: "",
		}
		if err := db.DB.Create(&newUser).Error; err != nil {
			http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
			return
		}
		user = newUser
	}

	sessionToken, err := CreateSession(user.UserID)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionToken",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}
