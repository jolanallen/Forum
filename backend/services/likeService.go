package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"

	"gorm.io/gorm"
)

func HasUserLikedPost(userID, postID uint64) (bool, error) {
	var like structs.PostLike
	err := db.DB.Where("userID = ? AND postID = ?", userID, postID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}


func AddLikeToPost(userID, postID uint64, post *structs.Post) error {
	newLike := structs.PostLike{
		UserID: userID,
		PostID: postID,
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	post.PostLike++
	return db.DB.Save(post).Error
}


func RemoveLikeFromPost(userID, postID uint64, post *structs.Post) error {
	if err := db.DB.Where("userID = ? AND postID = ?", userID, postID).Delete(&structs.PostLike{}).Error; err != nil {
		return err
	}
	if post.PostLike > 0 {
		post.PostLike--
	}
	return db.DB.Save(post).Error
}


func HasUserLikedComment(userID, commentID uint64) (bool, error) {
	var like structs.CommentLike
	err := db.DB.Where("userID = ? AND commentID = ?", userID, commentID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func AddLikeToComment(userID, commentID uint64, comment *structs.Comment) error {
	newLike := structs.CommentLike{
		UserID:    userID,
		CommentID: commentID,
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	comment.CommentLike++
	return db.DB.Save(comment).Error
}


func RemoveLikeFromComment(userID, commentID uint64, comment *structs.Comment) error {
	if err := db.DB.Where("userID = ? AND commentID = ?", userID, commentID).Delete(&structs.CommentLike{}).Error; err != nil {
		return err
	}
	if comment.CommentLike > 0 {
		comment.CommentLike--
	}
	return db.DB.Save(comment).Error
}
