package structs

type Like struct {
	UserID uint64 `gorm:"column:userID;primaryKey"`
	PostID uint64 `gorm:"column:postID;primaryKey"`
	Type   string `gorm:"size:255;not null"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
	Post Post `gorm:"foreignKey:PostID;references:PostID"`
}
