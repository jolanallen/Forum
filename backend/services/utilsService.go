package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var F = &structs.Forum{}

// RenderTemplate is a helper function to render an HTML template with the provided data.
// It looks for the template in the "web/templates" directory and handles errors during parsing and execution.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Construct the full path for the template file
	templatePath := filepath.Join("web/templates", tmpl)
	// Parse the template file
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}

	// Execute the template with the given data
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error displaying content", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

// ExtractIDFromURL extracts an ID from the URL path. 
// It assumes the ID is the third part of the URL (after splitting by "/").
func ExtractIDFromURL(path string) (uint64, error) {
	// Split the URL path by "/"
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, errors.New("Invalid URL, missing ID")
	}
	// Convert the ID part of the path from string to uint64
	id, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID: %v", err)
	}
	return id, nil
}

// GenerateToken creates a random token of 32 bytes and returns it as a hexadecimal string.
// It uses the crypto/rand package to generate a secure token.
func GenerateToken() string {
	// Generate 32 random bytes
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // If the token generation fails, panic.
	}
	// Return the token as a hexadecimal string
	return hex.EncodeToString(b)
}

// CreateUserSession creates a session for a user with a 24-hour expiration.
// It inserts a session record into the database and returns the session token.
func CreateUserSession(userID int64) (string, error) {
	// Generate a session token
	sessionToken := GenerateToken()
	// Set the expiration time to 24 hours from now
	expiration := time.Now().Add(24 * time.Hour)

	// Insert the session into the database
	query := `
		INSERT INTO sessionsUsers (userID, sessionToken, expiresAt, createdAt)
		VALUES (?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, userID, sessionToken, expiration, time.Now())
	if err != nil {
		return "", err
	}

	// Return the generated session token
	return sessionToken, nil
}

// CheckIfEmailExists checks if a user with the given email already exists in the database.
// If the user is found, it returns the user; otherwise, it returns an error.
func CheckIfEmailExists(email string) (*structs.User, error) {
	// Check for the user by email
	fmt.Println(email)
	user, err := GetUserByEmail(email)
	fmt.Println(user)
	if err != nil || user == nil {
		return nil, fmt.Errorf("Unknown user")
	}
	// Return the found user
	return user, nil
}

// CheckRegistrationForm validates the registration form fields. 
// It checks that all fields are filled in and that the password and confirmation match.
func CheckRegistrationForm(username, email, password, confirmPassword string) error {
	// Check if any field is empty
	if username == "" || email == "" || password == "" || confirmPassword == "" {
		return fmt.Errorf("All fields must be filled")
	}
	// Check if the password and confirmation match
	if password != confirmPassword {
		return fmt.Errorf("Passwords do not match")
	}
	return nil
}

// CheckPasswordHash compares the provided password with the stored hash.
// It returns true if the password matches the hash, otherwise false.
func CheckPasswordHash(password, hash string) bool {
	// Compare the password with the hashed value
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Return whether the passwords match (nil error indicates match)
	return err == nil
}

// HashPassword hashes the provided password using bcrypt.
// It returns the hashed password or an error if hashing fails.
func HashPassword(password string) (string, error) {
	// Hash the password with bcrypt's default cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
