package structs

import (
	"time"
)

type Admin struct {
	AdminID           uint64 `gorm:"column:adminID;primaryKey;autoIncrement"`
	AdminUsername     string `gorm:"column:adminUsername;unique;not null;size:255"`
	AdminPasswordHash string `gorm:"column:adminPasswordHash;not null"`
	AdminEmail        string `gorm:"column:adminEmail;unique;not null;size:255"`
	SessionID         uint64 `gorm:"column:sessionID"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID;constraint:OnDelete:SET NULL"`
}
