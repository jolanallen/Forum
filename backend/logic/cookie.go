package logic

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)


var db *sql.DB


func generateSessionToken(userID string) string {
	return userID + "_token"
}


func StoreSession(w http.ResponseWriter, userID, role string, duration time.Duration) {
	sessionToken := generateSessionToken(userID)
	expiresAt := time.Now().Add(duration)


	query := `INSERT INTO sessions (user_id, role, token, expires_at)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			token = VALUES(token),
			expires_at = VALUES(expires_at)
`

	_, err := db.Exec(query, userID, role, sessionToken, expiresAt)
	if err != nil {
		log.Println("error storing session in DB", err)
		return
	}

	cookie := &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,	
		Expires: expiresAt,
		HttpOnly: true,
		Secure: true,
		SameSite: http.SameSiteDefaultMode,
		Path: "/",
	}

	http.SetCookie(w, cookie)
}


func CreateAdminCookie(w http.ResponseWriter, adminID string) {
	StoreSession(w, adminID, "admin", 12*time.Hour)
}


func CreateUserCookie(w http.ResponseWriter, userID string) {
	StoreSession(w, userID, "user", 24*time.Hour)
}


func CreateGuestCookie(w http.ResponseWriter, guestID string) {
	StoreSession(w, guestID, "guest", 1*time.Hour)
}


