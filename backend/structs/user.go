package structs

import (
	"time"
)

type User struct {
	UserID             uint64 `gorm:"column:user_id;primaryKey;autoIncrement"`
	UserUsername       string `gorm:"column:user_username;unique;not null;size:255"`
	UserEmail          string `gorm:"column:user_email;unique;not null;size:255"`
	UserPasswordHash   string `gorm:"column:user_password_hash;not null"`
	UserProfilePicture string `gorm:"column:user_profile_picture;size:255"`

	SessionID uint64 `gorm:"column:session_id"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID;constraint:OnDelete:SET NULL"`
}
