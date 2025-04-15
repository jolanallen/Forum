package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"

	"gorm.io/gorm"
)

func HasUserLikedPost(userID, postID uint64) (bool, error) {
	var like structs.Like
	err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, postID, "Post").First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func AddLikeToPost(userID, postID uint64, post *structs.Post) error {
	newLike := structs.Like{
		UserID: userID,
		PostID: &postID,
		Type:   "Post",
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	post.PostLike++
	return db.DB.Save(post).Error
}

func RemoveLikeFromPost(userID, postID uint64, post *structs.Post) error {
	if err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, postID, "Post").Delete(&structs.Like{}).Error; err != nil {
		return err
	}
	if post.PostLike > 0 {
		post.PostLike--
	}
	return db.DB.Save(post).Error
}

func HasUserLikedComment(userID, commentID uint64) (bool, error) {
	var like structs.Like
	err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, commentID, "comment").First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func AddLikeToComment(userID, commentID uint64, comment *structs.Comment) error {
	newLike := structs.Like{
		UserID: userID,
		PostID: &commentID,
		Type:   "comment",
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	comment.CommentLike++
	return db.DB.Save(comment).Error
}

func RemoveLikeFromComment(userID, commentID uint64, comment *structs.Comment) error {
	if err := db.DB.Where("userID = ? AND commentID = ? AND type = ?", userID, commentID, "comment").Delete(&structs.Like{}).Error; err != nil {
		return err
	}
	if comment.CommentLike > 0 {
		comment.CommentLike--
	}
	return db.DB.Save(comment).Error
}