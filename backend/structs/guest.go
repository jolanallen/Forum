package structs

import (
	"time"
)

// Guest represents a guest user in the forum. It is used to track information about users who are not registered or logged in.
// The struct contains information on when the guest was created and their last visit time.
// It helps verify if the guest's information is correctly transitioned into a user when they create an account or log in.
type Guest struct {
	// GuestID is a unique identifier for the guest.
	GuestID uint64
	
	// GuestCreatedAt stores the timestamp when the guest was first created or visited.
	GuestCreatedAt time.Time
	
	// GuestLastVisitedAt stores the timestamp of the guest's last visit. It is optional and may be nil.
	GuestLastVisitedAt *time.Time
}
