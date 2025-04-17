package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

// GetUserIDFromSession retrieves the user ID from the session cookie.
func GetUserIDFromSession(r *http.Request) uint64 {
	// Get the session cookie
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0 // Return 0 if no session cookie is found
	}

	var userID uint64
	sessionType := cookie.Value[:5]

	// Extract userID based on session type
	switch sessionType {
	case "user_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "admin_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "guest_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	default:
		return 0 // Return 0 for an invalid session type
	}

	return userID
}

// GetUserByID retrieves a user by their user ID.
func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
	// Query to fetch user details by userID
	query := `
		SELECT userID, userEmail, userUsername, userPasswordHash, userProfilePicture, createdAt 
		FROM users 
		WHERE userID = ?
	`
	err := db.DB.QueryRow(query, userID).Scan(
		&user.UserID,
		&user.UserEmail,
		&user.UserUsername,
		&user.UserPasswordHash,
		&user.UserProfilePicture,
		&user.UserCreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil if no user is found
			return nil, nil
		}
		// Return error if any other issue occurs
		return nil, err
	}
	return &user, nil
}

// GetPostsByCategory retrieves all posts under a specific category.
func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post
	// Sub-query to get the categoryID by category name
	subQuery := "SELECT categoryID FROM categories WHERE LOWER(categoryName) = LOWER(?)"
	// Query to fetch posts by categoryID
	query := `
		SELECT postID, userID, categoryID, postKey, imageID, postComment, postLike, createdAt 
		FROM posts 
		WHERE categoryID = (` + subQuery + `)
	`

	rows, err := db.DB.Query(query, category)
	if err != nil {
		return nil, err // Return error if query fails
	}
	defer rows.Close()

	// Iterate over the query results and append them to the posts slice
	for rows.Next() {
		var post structs.Post
		err := rows.Scan(
			&post.PostID,
			&post.UserID,
			&post.CategoryID,
			&post.PostKey,
			&post.ImageID,
			&post.PostLike,
			&post.PostCreatedAt,
		)
		if err != nil {
			return nil, err // Return error if scanning fails
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// GetUserByEmail retrieves a user by their email address.
func GetUserByEmail(email string) (*structs.User, error) {
	fmt.Println(email)
	var user structs.User
	// Query to fetch user details by email
	query := `
		SELECT userID, userEmail, userUsername, userPasswordHash, userProfilePicture, createdAt 
		FROM users 
		WHERE userEmail = ?
	`
	err := db.DB.QueryRow(query, email).Scan(
		&user.UserID,
		&user.UserEmail,
		&user.UserUsername,
		&user.UserPasswordHash,
		&user.UserProfilePicture,
		&user.UserCreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return nil if no user is found
			return nil, nil
		}
		// Return error if any other issue occurs
		return nil, err
	}
	fmt.Println("blablaba")
	return &user, nil
}

// GetAdminDashboardData retrieves data for the admin dashboard, including counts of users, posts, comments, and guests.
func GetAdminDashboardData(adminID uint64) (*structs.AdminDashboardData, error) {
	var data structs.AdminDashboardData
	data.AdminID = adminID
	data.GeneratedAt = time.Now()

	var count int

	// Count total number of users
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return nil, err // Return error if query fails
	}
	data.TotalUsers = uint64(count)

	// Count total number of posts
	err = db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return nil, err // Return error if query fails
	}
	data.TotalPosts = uint64(count)

	// Count total number of comments
	err = db.DB.QueryRow("SELECT COUNT(*) FROM comments").Scan(&count)
	if err != nil {
		return nil, err // Return error if query fails
	}
	data.TotalComments = uint64(count)

	// Count total number of guests
	err = db.DB.QueryRow("SELECT COUNT(*) FROM guests").Scan(&count)
	if err != nil {
		return nil, err // Return error if query fails
	}
	data.TotalGuests = uint64(count)

	return &data, nil
}
