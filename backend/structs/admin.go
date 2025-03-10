package structs

import (
	"time"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"size:100;not null;unique"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"size:255;not null;unique"`
	CookieID     string    `gorm:"size:255"`
	AdminKey     string    `gorm:"size:255"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
