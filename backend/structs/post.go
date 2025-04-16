package structs

import (
	"time"
)

// Post represents a forum post made by a user.
// It contains details about the post, such as its unique ID, category, associated image, content, likes, creation date, and user information.
type Post struct {
	// PostID is the unique identifier for the post.
	PostID uint64
	
	// CategoryID is the ID of the category the post belongs to.
	CategoryID uint64
	
	// PostKey is a unique key used to reference the post.
	PostKey string
	
	// ImageID is the ID of the image associated with the post. It is nullable.
	ImageID *uint64
	
	// PostComment contains the content or text of the post.
	PostComment string
	
	// PostLike represents the number of likes the post has received.
	PostLike int
	
	// PostCreatedAt is the timestamp indicating when the post was created.
	PostCreatedAt time.Time
	
	// UserID is the ID of the user who created the post.
	UserID uint64
	
	// UserUsername is the username of the user who created the post.
	UserUsername string // Added to store the username of the user
}

