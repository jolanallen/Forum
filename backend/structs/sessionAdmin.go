package structs

import "time"

// SessionAdmin represents an admin session in the system.
// It stores information about the session token and its related details for an admin user.
type SessionAdmin struct {
	// ASessionID is the unique identifier for the admin session.
	ASessionID uint64
	
	// AAdminID is the ID of the admin associated with this session.
	AAdminID uint64
	
	// ASessionToken is the token generated for the admin's session.
	ASessionToken string
	
	// AExpiresAt is the timestamp indicating when the admin session will expire.
	AExpiresAt time.Time
	
	// ACreatedAt is the timestamp indicating when the admin session was created.
	ACreatedAt time.Time
}
