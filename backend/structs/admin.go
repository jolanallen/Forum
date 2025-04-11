package structs

import (
	"time"
)

type Admin struct {
	AdminID           uint64 `gorm:"column:admin_id;primaryKey;autoIncrement"`
	AdminUsername     string `gorm:"column:admin_username;unique;not null;size:255"`
	AdminPasswordHash string `gorm:"column:admin_password_hash;not null"`
	AdminEmail        string `gorm:"column:admin_email;unique;not null;size:255"`
	SessionID         uint64 `gorm:"column:session_id"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID;constraint:OnDelete:SET NULL"`
}
