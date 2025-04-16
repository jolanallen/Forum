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
	query := "SELECT postID, userID, categoriesID, content, postLike, status, visible, createdAt FROM posts WHERE postID = ?"
	err := db.DB.QueryRow(query, postID).Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.PostLike, &post.PostCreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, nil // Aucun post trouvé
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
	fmt.Sscanf(cookie.Value, "%d", &userID)
	return userID
}

// Récupère un utilisateur par son ID
func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
	query := "SELECT userID, userEmail, userName, userPassword, userRole FROM users WHERE userID = ?"
	err := db.DB.QueryRow(query, userID).Scan(&user.UserID, &user.UserEmail, &user.UserUsername, &user.UserPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Aucun utilisateur trouvé
		}
		return nil, err
	}
	return &user, nil
}

// Donne la structure de tous les posts en fonction de la catégorie
func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post
	subQuery := "SELECT categoriesID FROM categories WHERE LOWER(categoryName) = LOWER(?)"
	rows, err := db.DB.Query("SELECT postID, userID, categoryID, content, postLike, status, visible, createdAt FROM posts WHERE categoriesID = ("+subQuery+")", category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.PostID, &post.UserID, &post.CategoryID, &post.PostLike, &post.PostCreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Récupère un commentaire par son ID
func GetCommentByID(commentID uint64) (structs.Comment, error) {
	var comment structs.Comment
	query := "SELECT commentID, userID, postID, content, status, visible, createdAt FROM comments WHERE commentID = ?"
	err := db.DB.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.CommentContent, &comment.CommentStatus, &comment.CommentVisible, &comment.CommentCreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return comment, nil // Aucun commentaire trouvé
		}
		return comment, err
	}
	return comment, nil
}

// Récupère un utilisateur par son email
func GetUserByEmail(email string) (*structs.User, error) {
	var user structs.User
	query := "SELECT userID, userEmail, userName, userPassword, userRole FROM users WHERE userEmail = ?"
	err := db.DB.QueryRow(query, email).Scan(&user.UserID, &user.UserEmail, &user.UserUsername, &user.UserPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Aucun utilisateur trouvé
		}
		return nil, err
	}
	return &user, nil
}

// Récupère les données du tableau de bord administrateur
func GetAdminDashboardData(adminID uint64) (*structs.AdminDashboardData, error) {
	var data structs.AdminDashboardData
	data.AdminID = adminID
	data.GeneratedAt = time.Now()

	var count int

	// Total des utilisateurs
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalUsers = uint64(count)

	// Total des posts
	err = db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalPosts = uint64(count)

	// Total des commentaires
	err = db.DB.QueryRow("SELECT COUNT(*) FROM comments").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalComments = uint64(count)

	// Total des invités
	err = db.DB.QueryRow("SELECT COUNT(*) FROM guests").Scan(&count)
	if err != nil {
		return nil, err
	}
	data.TotalGuests = uint64(count)

	return &data, nil
}
