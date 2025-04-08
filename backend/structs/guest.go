package structs

import (
	"time"
)

type Guest struct {
	ID             uint      `gorm:"primaryKey"`
	CookieID       *int      `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	LastVisitedAt  *time.Time
}
