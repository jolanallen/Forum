package structs

type PostLike struct {
	UserID uint64 `gorm:"column:userID;primaryKey;not null"`
	PostID uint64 `gorm:"column:postID;primaryKey;not null"`

	// Optionnel : seulement si tu veux charger les infos du user
	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}