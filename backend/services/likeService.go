package services

import (
	"Forum/backend/db"
	"errors"
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

// AddLikeToPost adds a like to a post and updates the post's like count.
func AddLikeToPost(userID, postID uint64) error {
	// Insert a record into the postsLikes table to represent the like
	insert := `INSERT INTO postsLikes (userID, postID) VALUES (?, ?)`
	_, err := db.DB.Exec(insert, userID, postID)
	if err != nil {
		return err // Return error if insertion fails
	}

	// Update the like count for the post
	update := `UPDATE posts SET postLike = postLike + 1 WHERE postID = ?`
	_, err = db.DB.Exec(update, postID)
	return err // Return any error from updating the post like count
}

// RemoveLikeFromPost removes a like from a post and updates the post's like count.
func RemoveLikeFromPost(userID, postID uint64) error {
	// Delete the record from postsLikes to remove the like
	delete := `DELETE FROM postLikes WHERE userID = ? AND postID = ?`
	res, err := db.DB.Exec(delete, userID, postID)
	if err != nil {
		return err // Return error if deletion fails
	}
	// Check the number of affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err // Return error if the rowsAffected check fails
	}
	// Return an error if no rows were affected (like not found)
	if rowsAffected == 0 {
		return errors.New("like not found")
	}

	// Update the like count for the post, ensuring it doesn't go negative
	update := `UPDATE posts SET postLike = postLike - 1 WHERE postID = ? AND postLike > 0`
	_, err = db.DB.Exec(update, postID)
	return err // Return any error from updating the post like count
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

// AddLikeToComment adds a like to a comment and updates the comment's like count.
func AddLikeToComment(userID, commentID uint64) error {
	// Insert a record into the commentsLikes table to represent the like
	insert := `INSERT INTO commentsLikes (userID, commentID) VALUES (?, ?)`
	_, err := db.DB.Exec(insert, userID, commentID)
	if err != nil {
		return err // Return error if insertion fails
	}

	// Update the like count for the comment
	update := `UPDATE comments SET commentLike = commentLike + 1 WHERE commentID = ?`
	_, err = db.DB.Exec(update, commentID)
	return err // Return any error from updating the comment like count
}

// RemoveLikeFromComment removes a like from a comment and updates the comment's like count.
func RemoveLikeFromComment(userID, commentID uint64) error {
	// Delete the record from commentsLikes to remove the like
	delete := `DELETE FROM commentLikes WHERE userID = ? AND commentID = ?`
	res, err := db.DB.Exec(delete, userID, commentID)
	if err != nil {
		return err // Return error if deletion fails
	}
	// Check the number of affected rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err // Return error if the rowsAffected check fails
	}
	// Return an error if no rows were affected (like not found)
	if rowsAffected == 0 {
		return errors.New("like not found")
	}

	// Update the like count for the comment, ensuring it doesn't go negative
	update := `UPDATE comments SET commentLike = commentLike - 1 WHERE commentID = ? AND commentLike > 0`
	_, err = db.DB.Exec(update, commentID)
	return err // Return any error from updating the comment like count
}
