package structs

import (
	"time"
)

type Session struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	SessionToken string    `gorm:"size:255;not null;unique"`
	ExpiresAt   time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	User        User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
