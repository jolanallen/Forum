package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// UserCreatePost handles creating a new post by a user
// It processes the form data and uploads the image if necessary
func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	// If the request is not a POST, render the post creation form
	if r.Method != http.MethodPost {
		var categories []structs.Category
		rows, err := db.DB.Query("SELECT categoryID, categoryName FROM categories")
		if err != nil {
			http.Error(w, "Error fetching categories", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Populate categories
		for rows.Next() {
			var category structs.Category
			if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
				http.Error(w, "Error processing categories", http.StatusInternalServerError)
				return
			}
			categories = append(categories, category)
		}

		// Render the post creation page with available categories
		services.RenderTemplate(w, "create_post.html", struct {
			Categories []structs.Category
		}{
			Categories: categories,
		})
		return
	}

	// Extract userID from session and generate a unique post key
	userID := r.Context().Value("userID").(uint64)
	postKey := uuid.New().String()

	// Parse form values
	content, categoryID, err := services.ParseFormValues(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Handle image upload
	imageID, err := services.HandleImageUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new post into the database
	query := "INSERT INTO posts (postKey, postComment, userID, imageID, categoryID) VALUES (?, ?, ?, ?, ?)"
	_, err = db.DB.Exec(query, postKey, content, userID, imageID, categoryID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating post: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect to the confirmation page
	http.Redirect(w, r, "/post-created", http.StatusSeeOther)
}

// ToggleLikeComment handles toggling a like on a comment
func ToggleLikeComment(w http.ResponseWriter, r *http.Request) {
	// Get userID from session and extract commentID from URL
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	commentIDStr := vars["id"]

	// Convert commentID to integer
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	// Check if the user has already liked the comment
	hasLiked, err := services.HasUserLikedComment(userID, commentID)
	if err != nil {
		http.Error(w, "Error checking like status", http.StatusInternalServerError)
		return
	}

	// Add or remove the like based on the current status
	if hasLiked {
		_, err := db.DB.Exec("DELETE FROM commentsLikes WHERE userID = ? AND commentID = ?", userID, commentID)
		if err != nil {
			http.Error(w, "Error removing like", http.StatusInternalServerError)
			return
		}
	} else {
		_, err := db.DB.Exec("INSERT INTO commentsLikes (userID, commentID) VALUES (?, ?)", userID, commentID)
		if err != nil {
			http.Error(w, "Error adding like", http.StatusInternalServerError)
			return
		}
	}

	// Redirect back to the forum page
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// ToggleLikePost handles toggling a like on a post
func ToggleLikePost(w http.ResponseWriter, r *http.Request) {
	// Get userID from session and extract postID from URL
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	postIDStr := vars["id"]

	// Convert postID to integer
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Check if the user has already liked the post
	hasLiked, err := services.HasUserLikedPost(userID, postID)
	if err != nil {
		http.Error(w, "Error checking like status", http.StatusInternalServerError)
		return
	}

	// Add or remove the like based on the current status
	if hasLiked {
		_, err = db.DB.Exec("DELETE FROM postsLikes WHERE userID = ? AND postID = ?", userID, postID)
		if err != nil {
			http.Error(w, "Error removing like", http.StatusInternalServerError)
			return
		}
	} else {
		_, err = db.DB.Exec("INSERT INTO postsLikes (userID, postID) VALUES (?, ?)", userID, postID)
		if err != nil {
			http.Error(w, "Error adding like", http.StatusInternalServerError)
			return
		}
	}

	// Redirect back to the forum page
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// UserEditProfile handles editing a user's profile
func UserEditProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)

	// If the method is GET, render the profile edit form
	if r.Method == http.MethodGet {
		query := "SELECT userID, userUsername, userEmail, userProfilePicture FROM users WHERE userID = ?"
		var user structs.User
		err := db.DB.QueryRow(query, userID).Scan(&user.UserID, &user.UserUsername, &user.UserEmail, &user.UserProfilePicture)
		if err != nil {
			http.Error(w, "Error retrieving profile information", http.StatusInternalServerError)
			return
		}
		// Render the profile edit page with user data
		services.RenderTemplate(w, "profile_edit.html", user)
	} else if r.Method == http.MethodPost {
		// If the method is POST, process the form data
		newUsername := r.FormValue("userUsername")
		newEmail := r.FormValue("userEmail")
		newPassword := r.FormValue("userPassword")

		// Get the current user from the database
		user, err := services.GetUserByID(userID)
		if err != nil {
			http.Error(w, "Error retrieving user", http.StatusInternalServerError)
			return
		}

		// Update user fields if provided
		if newUsername != "" {
			user.UserUsername = newUsername
		}
		if newEmail != "" {
			user.UserEmail = newEmail
		}
		if newPassword != "" {
			// Hash the new password
			hashedPassword, err := services.HashPassword(newPassword)
			if err != nil {
				http.Error(w, "Error hashing password", http.StatusInternalServerError)
				return
			}
			user.UserPasswordHash = hashedPassword
		}

		// Handle profile picture upload
		if file, _, err := r.FormFile("userProfilePicture"); err == nil {
			image, err := services.ValidateImage(file, nil)
			if err != nil {
				http.Error(w, "Error validating image: "+err.Error(), http.StatusBadRequest)
				return
			}
			user.UserProfilePicture = int64(image.ImageID)
		}

		// Update the user data in the database
		_, err = db.DB.Exec("UPDATE users SET userUsername = ?, userEmail = ?, userPasswordHash = ?, userProfilePicture = ? WHERE userID = ?",
			user.UserUsername, user.UserEmail, user.UserPasswordHash, user.UserProfilePicture, user.UserID)
		if err != nil {
			http.Error(w, "Error updating profile", http.StatusInternalServerError)
			return
		}

		// Redirect to the updated profile page
		http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
	}
}

// Logout handles logging the user out by clearing the session cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// UserProfile handles displaying a user's profile
func UserProfile(w http.ResponseWriter, r *http.Request) {
	// Extract the userID from the URL
	userIDStr := r.URL.Path[len("/user/"):]

	// Convert the userID from string to uint64
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch the user data from the database
	user, err := services.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Render the user's profile page
	services.RenderTemplate(w, "profile.html", user)
}
