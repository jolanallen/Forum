package structs

import "time"

// SessionUser represents a user session in the system.
// It stores information about the session token and its related details for a registered user.
type SessionUser struct {
	// USessionID is the unique identifier for the user session.
	USessionID uint64
	
	// UUserID is the ID of the user associated with this session.
	// It is a foreign key linking to the users table.
	UUserID uint64
	
	// USessionToken is the token generated for the user's session.
	USessionToken string
	
	// UExpiresAt is the timestamp indicating when the user session will expire.
	UExpiresAt time.Time
	
	// UCreatedAt is the timestamp indicating when the user session was created.
	UCreatedAt time.Time
	
	// The relationship with User is removed for native SQL handling.
}
