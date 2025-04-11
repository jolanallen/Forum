package structs

import (
	"time"
)

type User struct {
	UserID           uint64    `gorm:"column:userID;primaryKey;autoIncrement"`
	UserUsername     string    `gorm:"column:userUsername;unique;not null;size:255"`
	UserEmail        string    `gorm:"column:userEmail;unique;not null;size:255"`
	UserPasswordHash string    `gorm:"column:userPasswordHash;not null"`
	UserProfilePicture string  `gorm:"column:userProfilePicture;size:255"`

	SessionID        uint64    `gorm:"column:sessionID"`

	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
}
