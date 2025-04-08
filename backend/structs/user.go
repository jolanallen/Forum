package structs

import (
	"time"
)

type User struct {
	UsersID        uint64    `gorm:"column:users_id;primaryKey"`
	Username       string    `gorm:"size:255;not null;unique"`
	PasswordHash   string    `gorm:"type:varchar(255);not null"`
	ProfilePicture string    `gorm:"size:255"`
	CookiesID      *int      `gorm:"default:null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	Posts          []Post    `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
	Comments       []Comment `gorm:"foreignKey:UserID;references:UsersID"`
	Topics         []Topic   `gorm:"foreignKey:UserID;references:UsersID"`
}
