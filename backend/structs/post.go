package structs

import (
	"time"
)

type Post struct {
	PostID      uint64    `gorm:"column:postID;primaryKey;autoIncrement"`
	PostKey   string    `gorm:"column:postKey;unique;not null;size:255"`
	PostComment string    `gorm:"column:postComment;type:text"`
	PostLike    int       `gorm:"column:postLike;default:0"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UserID      uint64    `gorm:"column:userID"`

	ImageID *uint64 `gorm:"column:imageID"`
	Image   Image   `gorm:"foreignKey:ImageID;references:ImageID"`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}
