package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

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

		user, err := services.CheckIfEmailExists(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !services.CheckPasswordHash(password, user.UserPasswordHash) {
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}

		sessionToken, err := services.CreateUserSession(user.UserID)
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

		err := services.CheckRegistrationForm(username, email, password, confirmPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if password != confirmPassword {
			http.Error(w, "Les mots de passe ne correspondent pas", http.StatusBadRequest)
			return
		}

		existingUser, err := services.GetUserByEmail(email)
		if err != nil {
			http.Error(w, "Erreur lors de la recherche de l'email", http.StatusInternalServerError)
			return
		}
		if existingUser != nil {
			http.Error(w, "Un compte avec cet email existe déjà", http.StatusBadRequest)
			return
		}

		hashedPassword, err := services.HashPassword(password)
		if err != nil {
			log.Println("Erreur lors du hachage du mot de passe:", err)
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}
		var userProfileImageID uint64
		userProfileImageID, err = services.HandleImageUpload(r)
		if err != nil {
			log.Println("Erreur d'upload d'image, utilisation de l'image par défaut")
			defaultImage := structs.Image{
				Filename: "default.png",
				URL:      "/images/default.png",
			}

			if err := db.DB.Create(defaultImage).Error; err != nil {
				log.Println("Erreur lors de l'ajout de l'image par défaut:", err)
				http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
				return
			}
			userProfileImageID = defaultImage.ImageID
		}
		newUser := structs.User{
			UserUsername:       username,
			UserEmail:          email,
			UserPasswordHash:   hashedPassword,
			UserProfilePicture: userProfileImageID,
		}

		if err := CreateUser(&newUser); err != nil {
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

	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'échange du code : %v", err), http.StatusInternalServerError)
		return
	}

	client := oauth2Config.Client(r.Context(), token)

	oauth2Service, err := oauth2api.NewService(r.Context(), option.WithHTTPClient(client))
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

	sessionToken, err := services.CreateUserSession(user.UserID)
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

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
