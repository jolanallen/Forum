package structs

import (
	"time"
)

type Comment struct {
	CommentID uint64    `gorm:"column:commentID;primaryKey;autoIncrement"`
	UserID    uint64    `gorm:"column:userID"`
	PostID    uint64    `gorm:"column:postID"`
	Content   string    `gorm:"column:content;type:text"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	Status    string    `gorm:"column:status;size:255"`
	Visible   bool      `gorm:"column:visible"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
	Post Post `gorm:"foreignKey:PostID;references:PostID;constraint:OnDelete:CASCADE"`
}
