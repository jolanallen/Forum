package structs

import "time"

type Comment struct {
	CommentID       uint64    `gorm:"column:comment_id;primaryKey;autoIncrement"`
	UserID          uint64    `gorm:"column:userID"`
	TopicID         uint64    `gorm:"column:topicID"`
	PostID          uint64    `gorm:"column:postID"`
	Content         string    `gorm:"column:content;type:text"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	Status          string    `gorm:"column:status"`
	Visible         bool      `gorm:"column:visible"`
	CommentsLike    int       `gorm:"column:comments_like"`
	CommentsDislike int       `gorm:"column:comments_dislike"`

	User  User  `gorm:"foreignKey:UserID;references:UsersID"`
	Topic Topic `gorm:"foreignKey:TopicID;references:TopicsID"`
	Post  Post  `gorm:"foreignKey:PostID;references:PostID"`
}

