package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"database/sql"
	"net/http"
)

// HandleCommentActions handles different HTTP methods for comment actions: POST, PUT, DELETE.
func HandleCommentActions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		UserAddComment(w, r) // Call function to add a comment
	case http.MethodPut:
		UserEditComment(w, r) // Call function to edit an existing comment
	case http.MethodDelete:
		UserDeleteComment(w, r) // Call function to delete a comment
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) // Return an error if the method is not allowed
	}
}

// UserAddComment handles adding a new comment.
func UserAddComment(w http.ResponseWriter, r *http.Request) {
	// Extract the post ID from the URL
	postID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest) // Return an error if the ID is invalid
		return
	}

	// Get the user ID from the context
	userID := r.Context().Value("userID").(uint64)
	// Get the comment content from the form data
	content := r.FormValue("comment")

	// Create a new comment object
	comment := structs.Comment{
		UserID:         userID,
		PostID:         postID,
		CommentContent: content,
		CommentStatus:  "published",
		CommentVisible: true,
	}

	// Insert the comment into the database with an SQL query
	query := "INSERT INTO comments (userID, postID, content, status, visible) VALUES (?, ?, ?, ?, ?)"
	_, err = db.DB.Exec(query, comment.UserID, comment.PostID, comment.CommentContent, comment.CommentStatus, comment.CommentVisible)
	if err != nil {
		http.Error(w, "Error adding comment", http.StatusInternalServerError) // Return error if the insertion fails
		return
	}

	// Redirect the user after adding the comment
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

// UserEditComment handles editing an existing comment.
func UserEditComment(w http.ResponseWriter, r *http.Request) {
	// Extract the comment ID from the URL
	commentID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest) // Return error if the comment ID is invalid
		return
	}

	// Retrieve the comment from the database
	var comment structs.Comment
	query := "SELECT commentID, userID, postID, content, status, visible FROM comments WHERE commentID = ?"
	err = db.DB.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.CommentContent, &comment.CommentStatus, &comment.CommentVisible)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound) // Return error if the comment does not exist
			return
		}
		http.Error(w, "Error fetching comment", http.StatusInternalServerError) // Return error if fetching fails
		return
	}

	// Check if the user is authorized to edit the comment
	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden) // Return error if the user is not the comment owner
		return
	}

	// Get the updated comment content
	updatedContent := r.FormValue("comment")
	comment.CommentContent = updatedContent

	// Update the comment in the database
	updateQuery := "UPDATE comments SET content = ? WHERE commentID = ?"
	_, err = db.DB.Exec(updateQuery, comment.CommentContent, comment.CommentID)
	if err != nil {
		http.Error(w, "Error updating comment", http.StatusInternalServerError) // Return error if update fails
		return
	}

	// Redirect after updating the comment
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

// UserDeleteComment handles deleting an existing comment.
func UserDeleteComment(w http.ResponseWriter, r *http.Request) {
	// Extract the comment ID from the URL
	commentID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Comment ID", http.StatusBadRequest) // Return error if the comment ID is invalid
		return
	}

	// Retrieve the comment from the database
	var comment structs.Comment
	query := "SELECT commentID, userID, postID, content, status, visible FROM comments WHERE commentID = ?"
	err = db.DB.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.CommentContent, &comment.CommentStatus, &comment.CommentVisible)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound) // Return error if the comment does not exist
			return
		}
		http.Error(w, "Error fetching comment", http.StatusInternalServerError) // Return error if fetching fails
		return
	}

	// Check if the user is authorized to delete the comment
	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Unauthorized", http.StatusForbidden) // Return error if the user is not the comment owner
		return
	}

	// Delete the comment from the database
	deleteQuery := "DELETE FROM comments WHERE commentID = ?"
	_, err = db.DB.Exec(deleteQuery, comment.CommentID)
	if err != nil {
		http.Error(w, "Error deleting comment", http.StatusInternalServerError) // Return error if delete fails
		return
	}

	// Redirect after deleting the comment
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}
