package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AdminDashboard handles the display of the admin dashboard
func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	if adminID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	adminIDUint, ok := adminID.(uint64)
	if !ok {
		http.Error(w, "Invalid admin ID", http.StatusBadRequest)
		return
	}

	// Fetch admin dashboard data
	data, err := services.GetAdminDashboardData(adminIDUint)
	if err != nil {
		http.Error(w, "Failed to load dashboard data", http.StatusInternalServerError)
		return
	}

	// Load and execute the template
	tmpl, err := template.ParseFiles("templates/admin_dashboard.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

// AdminDeleteUser handles user deletion by an admin
func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("userID")
	var admin structs.Admin

	// Validate admin from the database
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			http.Error(w, "Access error", http.StatusInternalServerError)
		}
		return
	}

	// Extract user ID from URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]

	// Convert the user ID to uint64
	userIDToDelete, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Check if the user exists
	var user structs.User
	row = db.DB.QueryRow("SELECT userID, userEmail FROM users WHERE userID = ?", userIDToDelete)
	if err := row.Scan(&user.UserID, &user.UserEmail); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		}
		return
	}

	// Delete the user from the database
	_, err = db.DB.Exec("DELETE FROM users WHERE userID = ?", userIDToDelete)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Redirect to dashboard
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

// AdminDeleteComment handles comment deletion by an admin
func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin

	// Validate admin from database
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			http.Error(w, "Access error", http.StatusInternalServerError)
		}
		return
	}

	// Extract comment ID from URL
	vars := mux.Vars(r)
	commentIDStr := vars["id"]

	// Convert the comment ID to uint64
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	// Check if the comment exists
	var comment structs.Comment
	row = db.DB.QueryRow("SELECT commentID, content FROM comments WHERE commentID = ?", commentID)
	if err := row.Scan(&comment.CommentID, &comment.CommentContent); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving comment", http.StatusInternalServerError)
		}
		return
	}

	// Delete the comment from the database
	_, err = db.DB.Exec("DELETE FROM comments WHERE commentID = ?", commentID)
	if err != nil {
		http.Error(w, "Error deleting comment", http.StatusInternalServerError)
		return
	}

	// Redirect to dashboard
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

// AdminDeletePost handles post deletion by an admin
func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin

	// Validate admin from database
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			http.Error(w, "Access error", http.StatusInternalServerError)
		}
		return
	}

	// Extract post ID from URL
	vars := mux.Vars(r)
	postIDStr := vars["id"]

	// Convert the post ID to uint64
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Check if the post exists
	var post structs.Post
	row = db.DB.QueryRow("SELECT postID FROM posts WHERE postID = ?", postID)
	if err := row.Scan(&post.PostID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving post", http.StatusInternalServerError)
		}
		return
	}

	// Delete the post from the database
	_, err = db.DB.Exec("DELETE FROM posts WHERE postID = ?", postID)
	if err != nil {
		http.Error(w, "Error deleting post", http.StatusInternalServerError)
		return
	}

	// Redirect to dashboard
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
