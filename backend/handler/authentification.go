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

// Google OAuth2 configuration
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

// Handles user login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")
		fmt.Println(email)

		// Check if email exists in the database
		var user structs.User
		row := db.DB.QueryRow("SELECT userID, userPasswordHash FROM users WHERE userEmail = $1", email)
		if err := row.Scan(&user.UserID, &user.UserPasswordHash); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Email not found", http.StatusUnauthorized)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
			return
		}

		// Validate password
		if !services.CheckPasswordHash(password, user.UserPasswordHash) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Create user session
		sessionToken, err := services.CreateUserSession(user.UserID)
		if err != nil {
			http.Error(w, "Session creation failed", http.StatusInternalServerError)
			return
		}

		// Store session token in cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionToken",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		// Redirect to homepage
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		services.RenderTemplate(w, "auth/login.html", nil)
	}
}

// Handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display registration page
		services.RenderTemplate(w, "auth/register.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("userUsername")
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")
		confirmPassword := r.FormValue("confirm_password")

		// Validate form fields
		err := services.CheckRegistrationForm(username, email, password, confirmPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if password != confirmPassword {
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}

		// Check if email is already taken
		var existingUser struct {
			UserEmail string
		}
		row := db.DB.QueryRow("SELECT userEmail FROM users WHERE userEmail = $1", email)
		if err := row.Scan(&existingUser.UserEmail); err != nil {
			if err != sql.ErrNoRows {
				http.Error(w, "Email check error", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Account with this email already exists", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := services.HashPassword(password)
		if err != nil {
			log.Println("Password hashing failed:", err)
			http.Error(w, "Registration error", http.StatusInternalServerError)
			return
		}

		// Handle profile image upload
		var userProfileImageID uint64
		userProfileImageID, err = services.HandleImageUpload(r)
		if err != nil {
			log.Println("Image upload failed, using default image")
			defaultImage := structs.Image{
				Filename: "default.png",
				URL:      "/images/default.png",
			}

			// Insert default image into database
			if err := db.DB.QueryRow(`
				INSERT INTO images (filename, url) VALUES ($1, $2) RETURNING imageID`,
				defaultImage.Filename, defaultImage.URL).Scan(&userProfileImageID); err != nil {
				log.Println("Failed to insert default image:", err)
				http.Error(w, "Registration error", http.StatusInternalServerError)
				return
			}
		}

		// Insert new user into database
		_, err = db.DB.Exec(`
			INSERT INTO users (userUsername, userEmail, userPasswordHash, userProfilePicture)
			VALUES ($1, $2, $3, $4)`,
			username, email, hashedPassword, userProfileImageID)
		if err != nil {
			log.Println("User creation failed:", err)
			http.Error(w, "Registration error", http.StatusInternalServerError)
			return
		}

		// Redirect to homepage
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

// Initiates Google login flow
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

// Handles callback from Google after authentication
func GoogleRegister(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	// Exchange code for OAuth token
	token, err := oauth2Config.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code exchange failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Create an OAuth client
	client := oauth2Config.Client(r.Context(), token)
	oauth2Service, err := oauth2api.NewService(r.Context(), option.WithHTTPClient(client))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create OAuth2 service: %v", err), http.StatusInternalServerError)
		return
	}

	// Get user info from Google
	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user info: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if user already exists
	var user structs.User
	row := db.DB.QueryRow("SELECT userID, userUsername FROM users WHERE userEmail = $1", userInfo.Email)
	if err := row.Scan(&user.UserID, &user.UserUsername); err != nil {
		// If user doesn't exist, create one
		_, err := db.DB.Exec(`
			INSERT INTO users (userUsername, userEmail) 
			VALUES ($1, $2)`,
			userInfo.Name, userInfo.Email)
		if err != nil {
			http.Error(w, "User creation failed", http.StatusInternalServerError)
			return
		}

		// Retrieve newly created user
		row = db.DB.QueryRow("SELECT userID FROM users WHERE userEmail = $1", userInfo.Email)
		if err := row.Scan(&user.UserID); err != nil {
			http.Error(w, "Failed to fetch new user", http.StatusInternalServerError)
			return
		}
	}

	// Create session for user
	sessionToken, err := services.CreateUserSession(user.UserID)
	if err != nil {
		http.Error(w, "Session creation failed", http.StatusInternalServerError)
		return
	}

	// Store session token in cookie
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
