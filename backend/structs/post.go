package structs

import (
	"time"
)

type Post struct {
	ID           uint      `gorm:"primaryKey"`
	PostKey      string    `gorm:"size:255;not null;unique"`
	PostImage    []byte    `gorm:"type:blob"`
	PostComments string    `gorm:"type:text"`
	PostLikes    int       `gorm:"default:0"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UserID       uint      `gorm:"default:null"`
	User         User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
