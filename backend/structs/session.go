package structs

import (
	"time"
)

type Session struct {
	SessionID   uint64    `gorm:"column:session_id;primaryKey;autoIncrement"`
	SessionToken string   `gorm:"column:session_token;unique;not null"`
	ExpiresAt   time.Time `gorm:"column:expires_at;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}
