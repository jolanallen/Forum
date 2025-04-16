package structs

import "time"

type SessionUser struct {
	USessionID    uint64    // sessionID
	UUserID       uint64    // userID (clé étrangère vers users.userID)
	USessionToken string    // sessionToken
	UExpiresAt    time.Time // expires_at
	UCreatedAt    time.Time // created_at

	// On vire le lien vers User pour rester simple en SQL pur
}
