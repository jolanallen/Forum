package structs

import (
	"time"
)

type User struct {
	UserID            uint64         `gorm:"primaryKey;autoIncrement"`
	UserUsername      string         `gorm:"size:255;not null;unique"`
	UserEmail         string         `gorm:"size:255;not null;unique"`
	UserPasswordHash  string         `gorm:"size:255;not null"`
	UserProfilePicture uint64        `gorm:"not null"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`

	ProfileImage      Image          `gorm:"foreignKey:UserProfilePicture;references:ImageID"`
	Sessions          []SessionUser  `gorm:"foreignKey:UserID"`
}


