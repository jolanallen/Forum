package structs

type Like struct {
	UserID    uint64  `gorm:"column:userID;primaryKey"`
	PostID    *uint64 `gorm:"column:postID"`    // Optionnel
	CommentID *uint64 `gorm:"column:commentID"` // Optionnel
	Type      string  `gorm:"size:255;not null;primaryKey"`

	User    User    `gorm:"foreignKey:UserID;references:UserID"`
	Post    *Post   `gorm:"foreignKey:PostID;references:PostID"`
	Comment *Comment `gorm:"foreignKey:CommentID;references:CommentID"`
}
