package structs

import "time"

type Comment struct {
	ID              uint      `gorm:"primaryKey"`
	UserID          uint      `gorm:"not null"`
	TopicID         uint      `gorm:"default:null"`
	PostID          uint      `gorm:"default:null"`
	Content         string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	Status          string    `gorm:"size:255"`
	Visible         bool      `gorm:"default:true"`
	CommentsLike    int       `gorm:"default:0"`
	CommentsDislike int       `gorm:"default:0"`
	User            User      `gorm:"foreignKey:UserID;references:ID"`
	Topic           Topic     `gorm:"foreignKey:TopicID;references:ID"`
	Post            Post      `gorm:"foreignKey:PostID;references:ID"`
}
