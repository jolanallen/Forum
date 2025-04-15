package structs

type CommentLike struct {
	UserID    uint64 `gorm:"column:userID;primaryKey;not null"`
	CommentID uint64 `gorm:"column:commentID;primaryKey;not null"`

	User    User    `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Comment Comment `gorm:"foreignKey:CommentID;references:CommentID;constraint:OnDelete:CASCADE"`
}