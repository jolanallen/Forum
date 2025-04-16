package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
	"time"
	"database/sql"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// Configuration OAuth2 pour Google
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

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")
		fmt.Println(email)

		// Vérifier si l'email existe dans la base de données
		var user structs.User
		row := db.DB.QueryRow("SELECT id, user_password_hash FROM users WHERE user_email = $1", email)
		if err := row.Scan(&user.UserID, &user.UserPasswordHash); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Email non trouvé", http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			return
		}

		// Vérifier le mot de passe
		if !services.CheckPasswordHash(password, user.UserPasswordHash) {
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}

		// Créer une session utilisateur
		sessionToken, err := services.CreateUserSession(user.UserID)
		if err != nil {
			http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
			return
		}

		// Stocker le token de session dans un cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionToken",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		// Rediriger vers la page d'accueil
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		services.RenderTemplate(w, "auth/login.html", nil)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		services.RenderTemplate(w, "auth/register.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("userUsername")
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")
		confirmPassword := r.FormValue("confirm_password")

		// Vérification des champs du formulaire
		err := services.CheckRegistrationForm(username, email, password, confirmPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if password != confirmPassword {
			http.Error(w, "Les mots de passe ne correspondent pas", http.StatusBadRequest)
			return
		}

		// Vérifier si l'email est déjà pris
		var existingUser struct {
			UserEmail string
		}
		row := db.DB.QueryRow("SELECT user_email FROM users WHERE user_email = $1", email)
		if err := row.Scan(&existingUser.UserEmail); err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, "Erreur lors de la vérification de l'email", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Un compte avec cet email existe déjà", http.StatusBadRequest)
			return
		}

		// Hachage du mot de passe
		hashedPassword, err := services.HashPassword(password)
		if err != nil {
			log.Println("Erreur lors du hachage du mot de passe:", err)
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}

		// Gestion de l'image de profil
		var userProfileImageID uint64
		userProfileImageID, err = services.HandleImageUpload(r)
		if err != nil {
			log.Println("Erreur d'upload d'image, utilisation de l'image par défaut")
			defaultImage := structs.Image{
				Filename: "default.png",
				URL:      "/images/default.png",
			}

			// Insérer l'image par défaut
			if err := db.DB.QueryRow(`
				INSERT INTO images (filename, url) VALUES ($1, $2) RETURNING image_id`,
				defaultImage.Filename, defaultImage.URL).Scan(&userProfileImageID); err != nil {
				log.Println("Erreur lors de l'ajout de l'image par défaut:", err)
				http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
				return
			}
		}

		// Création d'un nouvel utilisateur
		_, err = db.DB.Exec(`
			INSERT INTO users (user_username, user_email, user_password_hash, user_profile_picture)
			VALUES ($1, $2, $3, $4)`,
			username, email, hashedPassword, userProfileImageID)
		if err != nil {
			log.Println("Erreur lors de la création de l'utilisateur:", err)
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func GoogleRegister(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code manquant", http.StatusBadRequest)
		return
	}

	// Échanger le code contre un token OAuth
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'échange du code : %v", err), http.StatusInternalServerError)
		return
	}

	// Créer un client OAuth avec le token
	client := oauth2Config.Client(r.Context(), token)
	oauth2Service, err := oauth2api.NewService(r.Context(), option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la création du service OAuth2 : %v", err), http.StatusInternalServerError)
		return
	}

	// Récupérer les informations utilisateur
	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération des infos utilisateur : %v", err), http.StatusInternalServerError)
		return
	}

	// Vérifier si l'utilisateur existe déjà
	var user structs.User
	row := db.DB.QueryRow("SELECT id, user_username FROM users WHERE user_email = $1", userInfo.Email)
	if err := row.Scan(&user.UserID, &user.UserUsername); err != nil {
		// Si l'utilisateur n'existe pas, créer un nouvel utilisateur
		_, err := db.DB.Exec(`
			INSERT INTO users (user_username, user_email) 
			VALUES ($1, $2)`,
			userInfo.Name, userInfo.Email)
		if err != nil {
			http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
			return
		}

		// Récupérer l'utilisateur nouvellement créé
		row = db.DB.QueryRow("SELECT id FROM users WHERE user_email = $1", userInfo.Email)
		if err := row.Scan(&user.UserID); err != nil {
			http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusInternalServerError)
			return
		}
	}

	// Créer une session pour l'utilisateur
	sessionToken, err := services.CreateUserSession(user.UserID)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
		return
	}

	// Stocker le token de session dans un cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionToken",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
