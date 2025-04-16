package structs

// PostLike represents a "like" action on a forum post by a user.
// It stores the user who liked the post and the post that was liked.
type PostLike struct {
	// UserID is the ID of the user who liked the post.
	UserID uint64
	
	// PostID is the ID of the post that was liked.
	PostID uint64
}
