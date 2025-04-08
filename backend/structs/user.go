package structs

import (
	"time"
)

type User struct {
	ID             uint      `gorm:"primaryKey"`
	Username       string    `gorm:"size:255;not null;unique"`
	PasswordHash   string    `gorm:"type:varchar(255);not null"`
	ProfilePicture string    `gorm:"size:255"`
	CookiesID      *int      `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	Posts          []Post    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Comments       []Comment `gorm:"foreignKey:UserID;references:ID"`
	Topics         []Topic   `gorm:"foreignKey:UserID;references:ID"`
}
