package structs

import "time"

type SessionGuest struct {
	GSessionID    uint64    // sessionID
	GGuestID      uint64    // guestID
	GSessionToken string    // sessionToken
	GExpiresAt    time.Time // expires_at
	GCreatedAt    time.Time // created_at

	// On enlève la relation directe avec Guest pour SQL natif
}
