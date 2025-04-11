package structs

import (
	"time"
)

type User struct {
	UsersID        uint64    `gorm:"column:users_id;primaryKey;autoIncrement"`
	Username       string    `gorm:"column:username;size:255;not null;unique"`
	PasswordHash   string    `gorm:"column:password_hash;type:varchar(255);not null"`
	ProfilePicture string    `gorm:"column:profile_picture;size:255"`
	CookiesID      *int      `gorm:"column:cookies_id;default:null"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`

	Posts    []Post    `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
	Comments []Comment `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
	Topics   []Topic   `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
}

