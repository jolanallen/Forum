package structs

import (
	"time"
)

type Session struct {
	SessionID    uint64    `gorm:"column:session_id;primaryKey;autoIncrement"`
	UserID       uint64    `gorm:"column:user_id;not null"`
	SessionToken string    `gorm:"column:session_token;size:255;unique;not null"`
	ExpiresAt    time.Time `gorm:"column:expires_at;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
}

