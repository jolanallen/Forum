package structs

import (
	"time"
)

type Session struct {
	SessionID    uint64    `gorm:"column:sessionID;primaryKey;autoIncrement"`
	UserID       uint64    `gorm:"column:userID;not null"`
	SessionToken string    `gorm:"column:sessionToken;unique;not null"`
	ExpiresAt    time.Time `gorm:"column:expires_at;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
}
