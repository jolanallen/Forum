package structs

import "time"

type Topic struct {
	TopicsID      uint64    `gorm:"column:topics_id;primaryKey"`      // ID du topic (type uint64 pour UNSIGNED)
	CategoryID    uint64    `gorm:"column:topics_categoryID"`         // ID de la catégorie (type uint64 pour UNSIGNED)
	UserID        uint64    `gorm:"column:topics_userID"`             // ID de l'utilisateur (type uint64 pour UNSIGNED)
	Title         string    `gorm:"column:topics_title"`              // Titre du topic
	Content       string    `gorm:"column:topics_content;type:text"`  // Contenu du topic
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"` // Date de création
	TopicsLike    int       `gorm:"column:topics_like"`               // Nombre de likes
	TopicsDislike int       `gorm:"column:topics_dislike"`            // Nombre de dislikes

	User     User     `gorm:"foreignKey:UserID;references:UsersID"`          // Relation avec User
	Category Category `gorm:"foreignKey:CategoryID;references:CategoriesID"` // Relation avec Category
}
