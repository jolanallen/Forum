package structs

import (
	"time"
)

type User struct {
	UserID             int64
	UserUsername       string
	UserEmail          string
	UserPasswordHash   string
	UserProfilePicture int64
	UserCreatedAt      time.Time
}
