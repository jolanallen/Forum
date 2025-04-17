package services

import (
	"Forum/backend/db"
)

// HasUserLikedPost checks if a user has already liked a post by its postID.
func HasUserLikedPost(userID, postID uint64) (bool, error) {
	var count int
	// Query to count the number of likes from a user on a specific post
	query := `SELECT COUNT(*) FROM postsLikes WHERE userID = ? AND postID = ?`
	err := db.DB.QueryRow(query, userID, postID).Scan(&count)
	if err != nil {
		return false, err // Return error if the query fails
	}
	// Return true if the user has liked the post, else false
	return count > 0, nil
}

// HasUserLikedComment checks if a user has already liked a comment by its commentID.
func HasUserLikedComment(userID, commentID uint64) (bool, error) {
	var count int
	// Query to count the number of likes from a user on a specific comment
	query := `SELECT COUNT(*) FROM commentsLikes WHERE userID = ? AND commentID = ?`
	err := db.DB.QueryRow(query, userID, commentID).Scan(&count)
	if err != nil {
		return false, err // Return error if the query fails
	}
	// Return true if the user has liked the comment, else false
	return count > 0, nil
}
