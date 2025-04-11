package structs

import "time"

type Topic struct {
	TopicsID      uint64    `gorm:"column:topics_id;primaryKey;autoIncrement"`
	CategoryID    uint64    `gorm:"column:topics_categoryID;not null"`
	UserID        uint64    `gorm:"column:topics_userID;not null"`
	Title         string    `gorm:"column:topics_title;size:255;not null"`
	Content       string    `gorm:"column:topics_content;type:text;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	TopicsLike    int       `gorm:"column:topics_like;default:0"`
	TopicsDislike int       `gorm:"column:topics_dislike;default:0"`

	User     User     `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
	Category Category `gorm:"foreignKey:CategoryID;references:CategoriesID;constraint:OnDelete:SET NULL"`
}

