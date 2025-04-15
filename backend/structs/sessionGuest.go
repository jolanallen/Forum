package structs

import "time"

type SessionGuest struct {
	SessionID    uint64    `gorm:"primaryKey;autoIncrement"`
	GuestID      uint64    `gorm:"not null"`
	SessionToken string    `gorm:"size:191;not null;unique"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Guest Guest `gorm:"foreignKey:GuestID;constraint:OnDelete:CASCADE;"`
}
