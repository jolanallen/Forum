package structs

import (
	"time"
)

// Comment represents a user's comment on a post in the system.
// It holds information regarding the content, status, visibility, and other attributes of the comment.
type Comment struct {
	// CommentID is the unique identifier for the comment.
	CommentID uint64
	
	// UserID represents the ID of the user who posted the comment.
	UserID uint64
	
	// PostID refers to the ID of the post to which the comment belongs.
	PostID uint64
	
	// CommentContent contains the actual text content of the comment.
	CommentContent string
	
	// CommentCreatedAt stores the timestamp when the comment was created.
	CommentCreatedAt time.Time
	
	// CommentStatus indicates the status of the comment (e.g., pending, approved, etc.).
	CommentStatus string
	
	// CommentVisible specifies whether the comment is visible to other users.
	CommentVisible bool
	
	// CommentLike represents the number of likes the comment has received.
	CommentLike int
}
