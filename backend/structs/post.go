package structs

import (
	"time"
)

type Post struct {
	PostID       uint64    `gorm:"column:post_id;primaryKey;autoIncrement"`
	PostKey      string    `gorm:"column:post_key;unique;not null;size:255"`
	PostImage    []byte    `gorm:"column:post_image"`
	PostComments string    `gorm:"column:post_comments;type:text"`
	PostLikes    int       `gorm:"column:post_likes;default:0"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UserID       uint64    `gorm:"column:user_id"`

	User User `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
}
