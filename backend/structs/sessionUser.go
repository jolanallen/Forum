package structs

import (
	"time"
)

type SessionUser struct {
	SessionID    uint64    `gorm:"primaryKey;autoIncrement"`
	UserID       uint64    `gorm:"not null"`
	SessionToken string    `gorm:"size:191;not null;unique"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}