package structs

import (
	"time"
)

type Post struct {
	PostID       uint64    `gorm:"column:post_id;primaryKey;autoIncrement"`
	PostKey      string    `gorm:"column:post_key;unique;not null;size:255"`
	PostComments string    `gorm:"column:post_comments;type:text"`
	PostLikes    int       `gorm:"column:post_likes;default:0"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UserID       uint64    `gorm:"column:user_id"`

	ImageID *uint64 `gorm:"column:image_id"`
	Image   Image   `gorm:"foreignKey:ImageID;references:ImageID"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}
