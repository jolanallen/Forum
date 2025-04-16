package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

// Récupère un post par son ID
func GetPostByID(postID uint64) (structs.Post, error) {
	var post structs.Post
	query := `
		SELECT postID, userID, categoryID, postKey, imageID, postComment, postLike, createdAt 
		FROM posts 
		WHERE postID = ?
	`
	err := db.DB.QueryRow(query, postID).Scan(
		&post.PostID,
		&post.UserID,
		&post.CategoryID,
		&post.PostKey,
		&post.ImageID,
		&post.PostLike,
		&post.PostCreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, nil
		}
		return post, err
	}
	return post, nil
}

// Récupère l'ID de l'utilisateur depuis le cookie de session
func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0
	}

	var userID uint64
	sessionType := cookie.Value[:5]

	switch sessionType {
	case "user_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "admin_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "guest_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	default:
		return 0
	}

	return userID
}

// Récupère un utilisateur par son ID
func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
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
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Récupère les posts par catégorie
func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post
	subQuery := "SELECT categoryID FROM categories WHERE LOWER(categoryName) = LOWER(?)"
	query := `
		SELECT postID, userID, categoryID, postKey, imageID, postComment, postLike, createdAt 
		FROM posts 
		WHERE categoryID = (` + subQuery + `)
	`

	rows, err := db.DB.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Récupère un commentaire par son ID
func GetCommentByID(commentID uint64) (structs.Comment, error) {
	var comment structs.Comment
	query := `
		SELECT commentID, userID, postID, content, status, visible, createdAt 
		FROM comments 
		WHERE commentID = ?
	`
	err := db.DB.QueryRow(query, commentID).Scan(
		&comment.CommentID,
		&comment.UserID,
		&comment.PostID,
		&comment.CommentContent,
		&comment.CommentStatus,
		&comment.CommentVisible,
		&comment.CommentCreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, nil
		}
		return comment, err
	}
	return comment, nil
}

// Récupère un utilisateur par son email
func GetUserByEmail(email string) (*structs.User, error) {
	fmt.Println(email)
	var user structs.User
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
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("blablaba")
	return &user, nil
}

// Récupère les données du dashboard admin
func GetAdminDashboardData(adminID uint64) (*structs.AdminDashboardData, error) {
	var data structs.AdminDashboardData
	data.AdminID = adminID
	data.GeneratedAt = time.Now()

	var count int

	err := db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalUsers = uint64(count)

	err = db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalPosts = uint64(count)

	err = db.DB.QueryRow("SELECT COUNT(*) FROM comments").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalComments = uint64(count)

	err = db.DB.QueryRow("SELECT COUNT(*) FROM guests").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalGuests = uint64(count)

	return &data, nil
}
