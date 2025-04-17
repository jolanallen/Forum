package structs

// CommentLike represents the relationship between a user and a comment they liked.
// It holds the user ID and the comment ID to track which users have liked which comments.
type CommentLike struct {
	// UserID represents the ID of the user who liked the comment.
	UserID uint64
	
	// CommentID represents the ID of the comment that was liked by the user.
	CommentID uint64
}
