package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
)
// /services
// func ToggleLikePost
// récuper le post à partir de l'ID
func GetPostByID(postID uint64) (structs.Post, error) {
	var post structs.Post
	err := db.DB.First(&post, postID).Error
	return post, err
}


// /services
// func GuestHome
// récuperer l'utilisateur à partir de l'ID
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
// /services
// func UserEditProfile
// récuperer l'utilisateur à partir de l'ID
func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
// /services
// func GuestHack; GuestProg; GuestNews
// récupere les posts à partir des categories
func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post
	err := db.DB.Where("categoriesName = ?", category).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
// /services
// func ToggleLikeComment
// récuperer le commentaire à partir de l'ID
func GetCommentByID(commentID uint64) (structs.Comment, error) {
	var comment structs.Comment
	err := db.DB.First(&comment, commentID).Error
	return comment, err
}
// /services
// func Register; CheckAdmin
// récuperer l'utilisateur à partir du mail
func GetUserByEmail(email string) (*structs.User, error) {
	var user structs.User
	result := db.DB.Where("userEmail = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}





//inutilisé
func GetUserIDFromUsername(username string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("userUsername = ?", username).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
//inutilisé
func GetUserIDFromEmail(email string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("userEmail = ?", email).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
//inutilisé
func GetUserIDFromCommentID(commentID uint64) (uint64, error) {
	var comment structs.Comment
	if err := db.DB.Where("commentID = ?", commentID).First(&comment).Error; err != nil {
		return 0, fmt.Errorf("commentaire introuvable")
	}
	return comment.UserID, nil
}
//inutilisé
func GetUserIDFromPostID(postID uint64) (uint64, error) {
	var post structs.Post
	if err := db.DB.Where("postID = ?", postID).First(&post).Error; err != nil {
		return 0, fmt.Errorf("post introuvable")
	}
	return post.UserID, nil
}
//inutilisé
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(users)
}
//inutilisé
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
}
//inutilisé
func getAllPost(w http.ResponseWriter, r *http.Request) {
	var posts []structs.Post
	db.DB.Preload("User").Preload("Comments").Find(&posts)
	// faut changer vers le bon template
	Templates.ExecuteTemplate(w, "home.html", posts)
}
