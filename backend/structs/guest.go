package structs

import (
	"time"
)

type Guest struct {
	GuestID       uint64    `gorm:"column:guest_id;primaryKey;autoIncrement"`
	SessionID     uint64    `gorm:"column:session_id"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	LastVisitedAt time.Time `gorm:"column:last_visited_at"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID;constraint:OnDelete:SET NULL"`
}
