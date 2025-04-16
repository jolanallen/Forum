package structs

import (
	"time"
)

type Comment struct {
	CommentID        uint64
	UserID           uint64
	PostID           uint64
	CommentContent   string
	CommentCreatedAt time.Time
	CommentStatus    string
	CommentVisible   bool
	CommentLike      int
}
