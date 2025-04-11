package structs

import (
	"time"
)

type Guest struct {
	GuestID       uint64    `gorm:"column:guestID;primaryKey;autoIncrement"`
	SessionID     uint64    `gorm:"column:sessionID"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	LastVisitedAt time.Time `gorm:"column:last_visited_at"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID;constraint:OnDelete:SET NULL"`
}
