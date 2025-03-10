package structs

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"size:100;not null;unique"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"size:255;not null;unique"`
	Cookies      string    `gorm:"size:255"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
