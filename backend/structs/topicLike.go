package structs

type TopicLike struct {
	ID      uint `gorm:"primaryKey"`
	TopicID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
	Topic   Topic `gorm:"foreignKey:TopicID;references:ID"`
	User    User  `gorm:"foreignKey:UserID;references:ID"`
}
