package structs

type PostLike struct {
	UserID   uint64 `gorm:"column:userID;primaryKey"`
	PostID   uint64 `gorm:"column:postID;primaryKey"`
	PostLike bool   `gorm:"column:postLike;default:false"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
	Post Post `gorm:"foreignKey:PostID;references:PostID"`
}
