package structs

type PostLike struct {
	UserID   uint64 `gorm:"column:user_id;primaryKey"`
	PostID   uint64 `gorm:"column:post_id;primaryKey"`
	PostLike bool   `gorm:"column:post_like;default:false"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
	Post Post `gorm:"foreignKey:PostID;references:PostID"`
}
