package structs

type TopicLike struct {
	ID      uint64 `gorm:"column:topicsLikes_id;primaryKey"` // ID de la relation (type uint64 pour UNSIGNED)
	TopicID uint64 `gorm:"column:topicID"`                   // ID du topic (type uint64 pour UNSIGNED)
	UserID  uint64 `gorm:"column:userID"`                    // ID de l'utilisateur (type uint64 pour UNSIGNED)

	Topic Topic `gorm:"foreignKey:TopicID;references:TopicsID"` // Relation avec Topic
	User  User  `gorm:"foreignKey:UserID;references:UsersID"`   // Relation avec User
}
