package structs

import (
	"time"
)

type Guest struct {
	GuestID       uint       `gorm:"column:guests_id;primaryKey;autoIncrement"`
	CookieID      *int       `gorm:"column:guests_cookie_id"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime"`
	LastVisitedAt *time.Time `gorm:"column:last_visited_at"`
}
