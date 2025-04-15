package structs

import (
	"time"
)

type Session struct {
	SessionID    uint64    `gorm:"column:sessionID;primaryKey;autoIncrement"` // sessionID en clé primaire
	UserID       uint64    `gorm:"column:userID;not null"`                    // userID lié à la table users
	SessionToken string    `gorm:"column:sessionToken;unique;not null"`       // sessionToken unique
	ExpiresAt    time.Time `gorm:"column:expires_at;not null"`                // expiration du token
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`          // date de création automatique

	// Relation avec la table Users
	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}
