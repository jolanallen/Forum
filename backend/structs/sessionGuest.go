package structs

import "time"

// SessionGuest represents a guest session in the system.
// It stores information about the session token and its related details for a guest user.
type SessionGuest struct {
	// GSessionID is the unique identifier for the guest session.
	GSessionID uint64
	
	// GGuestID is the ID of the guest associated with this session.
	GGuestID uint64
	
	// GSessionToken is the token generated for the guest's session.
	GSessionToken string
	
	// GExpiresAt is the timestamp indicating when the guest session will expire.
	GExpiresAt time.Time
	
	// GCreatedAt is the timestamp indicating when the guest session was created.
	GCreatedAt time.Time
	
	// The relationship with Guest is removed for native SQL handling.
}
