package structs

import (
	"time"
)

// donc ici on va changer pour mettrel'image pour la photo de profil (faudra changer la bdd aussi)
type User struct {
	UserID             uint64  `gorm:"column:userID;primaryKey;autoIncrement"`
	UserUsername       string  `gorm:"column:userUsername;unique;not null;size:255"`
	UserEmail          string  `gorm:"column:userEmail;unique;not null;size:255"`
	UserPasswordHash   string  `gorm:"column:userPasswordHash;not null"`
	UserProfilePicture *uint64 `gorm:"column:userProfilePicture" json:"imageID"`
	SessionID          uint64  `gorm:"column:sessionID"`

	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	ProfilePicture Image     `gorm:"foreignKey:UserProfilePicture;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
