package structs

import "time"

type SessionAdmin struct {
	SessionID    uint64    `gorm:"primaryKey;autoIncrement"`
	AdminID      uint64    `gorm:"not null"`
	SessionToken string    `gorm:"size:191;not null;unique"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Admin Admin `gorm:"foreignKey:AdminID;constraint:OnDelete:CASCADE;"`
}
