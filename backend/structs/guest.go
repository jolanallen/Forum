package structs

import (
	"time"
)

type Guest struct {
    ID            uint      `gorm:"primaryKey"`
    CookieID      string    `gorm:"size:255"`
    CreatedAt     time.Time `gorm:"autoCreateTime"`
    LastVisitedAt time.Time
}