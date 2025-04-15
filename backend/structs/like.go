package structs

type Like struct {
	UserID    uint64 `gorm:"column:userID;primaryKey"`
	PostID    uint64 `gorm:"column:postID;primaryKey"`
	CommentID uint64 `gorm:"column:commentID;primaryKey"` // Ajout de la clé étrangère pour Comment
	Type      string `gorm:"size:255;not null"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
	Post Post `gorm:"foreignKey:PostID;references:PostID"`
	Comment Comment `gorm:"foreignKey:CommentID;references:CommentID"` // Ajout de la relation avec Comment
}
