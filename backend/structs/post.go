package structs

import (
	"time"
)
//on va reprendre la façon de faire du post image, ce sera plus rapide
type Post struct {
	PostID      uint64    `gorm:"primaryKey;autoIncrement;column:postID" json:"postID"`
	CategoryID  uint64    `gorm:"column:categoriesID" json:"categoryID"`
	PostKey     string    `gorm:"unique;column:postKey" json:"postKey"`
	ImageID     *uint64   `gorm:"column:imageID" json:"imageID"`
	PostComment string    `gorm:"column:postComment" json:"postComment"`
	PostLike    int       `gorm:"default:0;column:postLike" json:"postLike"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UserID      uint64    `gorm:"column:userID" json:"userID"`
}
