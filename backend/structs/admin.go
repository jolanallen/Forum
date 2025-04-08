package structs

import (
	"time"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"size:255;not null;unique"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"size:255;not null;unique"`
	CookieID     *int      `gorm:"default:null"`
	AdminKey     string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
