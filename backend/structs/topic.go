package structs

import "time"

type Topic struct {
	ID            uint           `gorm:"primaryKey"`
	CategoryID    uint           `gorm:"not null"`
	UserID        uint           `gorm:"not null"`
	Title         string         `gorm:"size:255;not null"`
	Content       string         `gorm:"type:text"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	TopicsLike    int            `gorm:"default:0"`
	TopicsDislike int            `gorm:"default:0"`
	Category      Category       `gorm:"foreignKey:CategoryID;references:ID"`
	User          User           `gorm:"foreignKey:UserID;references:ID"`
	Comments      []Comment      `gorm:"foreignKey:TopicID;references:ID"`
	TopicLikes    []TopicLike    `gorm:"foreignKey:TopicID;references:ID"`
	TopicDislikes []TopicDislike `gorm:"foreignKey:TopicID;references:ID"`
}
