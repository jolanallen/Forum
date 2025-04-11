package structs

import (
	"time"
)

type Admin struct {
	AdminID      uint64    `gorm:"column:admin_id;primaryKey;autoIncrement"`
	Username     string    `gorm:"column:admin_username;size:255;not null;unique"`
	PasswordHash string    `gorm:"column:admin_password_hash;size:255;not null"`
	Email        string    `gorm:"column:admin_email;size:255;not null;unique"`
	CookieID     *int      `gorm:"column:admin_cookie_id;default:null"`
	AdminKey     string    `gorm:"column:admin_key;size:255;not null;unique"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

