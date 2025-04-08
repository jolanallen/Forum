package structs

import "time"

type Comment struct {
	CommentID       uint64    `gorm:"column:comment_id;primaryKey"`     // Utiliser uint64 pour la clé primaire
	UserID          uint64    `gorm:"column:userID"`                    // Doit correspondre à `users_id` dans la table users
	TopicID         uint64    `gorm:"column:topicID"`                   // Doit correspondre à `topics_id` dans la table topics
	PostID          uint64    `gorm:"column:postID"`                    // Correspond à la table posts
	Content         string    `gorm:"column:content;type:text"`         // Contenu du commentaire
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"` // Date de création
	Status          string    `gorm:"column:status"`                    // Statut du commentaire
	Visible         bool      `gorm:"column:visible"`                   // Visibilité du commentaire
	CommentsLike    int       `gorm:"column:comments_like"`             // Nombre de likes
	CommentsDislike int       `gorm:"column:comments_dislike"`          // Nombre de dislikes

	User  User  `gorm:"foreignKey:UserID;references:UsersID"`   // Relation avec User
	Topic Topic `gorm:"foreignKey:TopicID;references:TopicsID"` // Relation avec Topic
	Post  Post  `gorm:"foreignKey:PostID;references:PostID"`    // Relation avec Post
}
