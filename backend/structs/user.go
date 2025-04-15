package structs

import (
	"time"
)

type User struct {
	UserID             uint64  `gorm:"column:userID;primaryKey;autoIncrement"`       // userID en clé primaire auto-incrémentée
	UserUsername       string  `gorm:"column:userUsername;unique;not null;size:255"` // userUsername unique et non nul
	UserEmail          string  `gorm:"column:userEmail;unique;not null;size:255"`    // userEmail unique et non nul
	UserPasswordHash   string  `gorm:"column:userPasswordHash;not null"`             // hash du mot de passe
	UserProfilePicture uint64 `gorm:"column:userProfilePicture" json:"imageID"`     // Clé étrangère vers `images` (photo de profil)
	SessionID          uint64  `gorm:"column:sessionID"`                             // sessionID associé à l'utilisateur

	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"` // `created_at` automatique
	ProfilePicture Image     `gorm:"foreignKey:UserProfilePicture;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
