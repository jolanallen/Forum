package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var Templates *template.Template

var F = &structs.Forum{}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templatePath := filepath.Join("web/templates", tmpl)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		log.Println("Erreur parsing template:", err)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage", http.StatusInternalServerError)
		log.Println("Erreur exécution template:", err)
	}
}

func ExtractIDFromURL(path string) (uint64, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, errors.New("URL invalide")
	}
	return strconv.ParseUint(parts[2], 10, 64)
}
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

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
