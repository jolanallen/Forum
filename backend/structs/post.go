package structs

import (
	"time"
)

type Post struct {
	PostID        uint64
	CategoryID    uint64
	PostKey       string
	ImageID       *uint64 // nullable
	PostComment   string
	PostLike      int
	PostCreatedAt time.Time
	UserID        uint64
	UserUsername  string // Ajout du champ UserUsername pour stocker le nom d'utilisateur
}

