package structs

import (
	"time"
)

type Post struct {
	PostID      uint64    `gorm:"primaryKey;autoIncrement;column:postID" json:"postID"`
	CategoryID  uint64    `gorm:"column:categoryID;not null" json:"categoryID"`
	PostKey     string    `gorm:"unique;column:postKey;not null" json:"postKey"`
	ImageID     uint64    `gorm:"column:imageID;not null" json:"imageID"`
	PostComment string    `gorm:"column:postComment" json:"postComment"`
	PostLike    int       `gorm:"default:0;column:postLike" json:"postLike"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UserID      uint64    `gorm:"column:userID" json:"userID"`
	Preview     string    `gorm:"-"`  // Ce champ ne sera pas mappé à la base de données

	// Relations
	Comments []Comment `gorm:"foreignKey:PostID;references:PostID;constraint:OnDelete:CASCADE" json:"comments"`
	Image    *Image    `gorm:"foreignKey:ImageID;references:ImageID" json:"image"`
	User     User      `gorm:"foreignKey:UserID;references:UserID" json:"user"`
}

