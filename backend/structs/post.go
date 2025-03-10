package structs

import (
	"time"
)

type Post struct {
    ID          uint      `gorm:"primaryKey"`
    PostKey     string    `gorm:"size:255;unique;not null"`
    Image       Image     `gorm:"embedded"`
    Commentaire string    `gorm:"type:text"`
    Likes       int       `gorm:"default:0"`
    Date        time.Time `gorm:"autoCreateTime"`
    CreatorID   uint      `gorm:"not null"`
    UserID      uint      `gorm:"not null"`
}