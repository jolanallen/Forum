package structs

type TopicLike struct {
	ID      uint64 `gorm:"column:topicsLikes_id;primaryKey;autoIncrement"`
	TopicID uint64 `gorm:"column:topicID;not null"`
	UserID  uint64 `gorm:"column:userID;not null"`

	Topic Topic `gorm:"foreignKey:TopicID;references:TopicsID;constraint:OnDelete:CASCADE"`
	User  User  `gorm:"foreignKey:UserID;references:UsersID;constraint:OnDelete:CASCADE"`
}

