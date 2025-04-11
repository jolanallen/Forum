package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
)

func GetPostByID(postID uint64) (structs.Post, error) {
	var post structs.Post
	err := db.DB.First(&post, postID).Error
	return post, err
}

func GetUserIDFromUsername(username string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("userUsername = ?", username).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromEmail(email string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("userEmail = ?", email).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromCommentID(commentID uint64) (uint64, error) {
	var comment structs.Comment
	if err := db.DB.Where("commentID = ?", commentID).First(&comment).Error; err != nil {
		return 0, fmt.Errorf("commentaire introuvable")
	}
	return comment.UserID, nil
}
func GetUserIDFromPostID(postID uint64) (uint64, error) {
	var post structs.Post
	if err := db.DB.Where("postID = ?", postID).First(&post).Error; err != nil {
		return 0, fmt.Errorf("post introuvable")
	}
	return post.UserID, nil
}

// le pb c'est qu'il prend par rapport à la session de son nav et pas de notre bdd
func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0
	}
	var userID uint64
	fmt.Sscanf(cookie.Value, "%d", &userID)
	return userID
}


func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post
	err := db.DB.Where("categoriesName = ?", category).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}









// Récupérer tous les users mais rien avec le front
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(users)
}

// Récupérer tous les admins
// ça ne return rien, c'est juste pour la prog
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
}

// fonction qui renvoie tout les posts de la base de données
func getAllPost(w http.ResponseWriter, r *http.Request) {
	var posts []structs.Post
	db.DB.Preload("User").Preload("Comments").Find(&posts)
	// faut changer vers le bon template
	Templates.ExecuteTemplate(w, "home.html", posts)
}