package utils

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
)

func GetUserIDFromUsername(username string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("users_username = ?", username).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromEmail(email string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromCommentID(commentID uint64) (uint64, error) {
	var comment structs.Comment
	if err := db.DB.Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		return 0, fmt.Errorf("commentaire introuvable")
	}
	return comment.UserID, nil
}
func GetUserIDFromPostID(postID uint64) (uint64, error) {
	var post structs.Post
	if err := db.DB.Where("post_id = ?", postID).First(&post).Error; err != nil {
		return 0, fmt.Errorf("post introuvable")
	}
	return post.UserID, nil
}
