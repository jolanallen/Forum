package structs

import "time"

// User represents a registered user in the system.
// It stores the user's personal information and account details.
type User struct {
	// UserID is the unique identifier for the user.
	UserID int64
	
	// UserUsername is the username chosen by the user.
	UserUsername string
	
	// UserEmail is the email address associated with the user.
	UserEmail string
	
	// UserPasswordHash is the hashed version of the user's password for authentication.
	UserPasswordHash string
	
	// UserProfilePicture is the ID of the user's profile picture.
	UserProfilePicture int64
	
	// UserCreatedAt is the timestamp when the user's account was created.
	UserCreatedAt time.Time
}
