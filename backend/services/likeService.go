package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"net/http"

	"gorm.io/gorm"
)
//
//
//
func ToggleLikePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	postID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	post, err := GetPostByID(postID)
	if err != nil {
		http.Error(w, "Post introuvable", http.StatusNotFound)
		return
	}

	hasLiked, err := hasUserLikedPost(userID, postID)
	if err != nil {
		http.Error(w, "Erreur lors de la vérification du like", http.StatusInternalServerError)
		return
	}

	if hasLiked {
		if err := removeLikeFromPost(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := addLikeToPost(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func hasUserLikedPost(userID, postID uint64) (bool, error) {
	var like structs.Like
	err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, postID, "Post").First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func addLikeToPost(userID, postID uint64, post *structs.Post) error {
	newLike := structs.Like{
		UserID: userID,
		PostID: postID,
		Type:   "Post",
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	post.PostLike++
	return db.DB.Save(post).Error
}

func removeLikeFromPost(userID, postID uint64, post *structs.Post) error {
	if err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, postID,"Post").Delete(&structs.Like{}).Error; err != nil {
		return err
	}
	if post.PostLike > 0 {
		post.PostLike--
	}
	return db.DB.Save(post).Error
}

func ToggleLikeComment(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	commentID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	comment, err := GetCommentByID(commentID)
	if err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	hasLiked, err := hasUserLikedComment(userID, commentID)
	if err != nil {
		http.Error(w, "Erreur lors de la vérification du like", http.StatusInternalServerError)
		return
	}

	if hasLiked {
		if err := removeLikeFromComment(userID, commentID, &comment); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := addLikeToComment(userID, commentID, &comment); err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func hasUserLikedComment(userID, commentID uint64) (bool, error) {
	var like structs.Like
	err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, commentID, "comment").First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func addLikeToComment(userID, commentID uint64, comment *structs.Comment) error {
	newLike := structs.Like{
		UserID: userID,
		PostID: commentID,
		Type:   "comment",
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	comment.CommentLike++
	return db.DB.Save(comment).Error
}

func removeLikeFromComment(userID, commentID uint64, comment *structs.Comment) error {
	if err := db.DB.Where("userID = ? AND postID = ? AND type = ?", userID, commentID, "comment").Delete(&structs.Like{}).Error; err != nil {
		return err
	}

	if comment.CommentLike > 0 {
		comment.CommentLike--
	}
	return db.DB.Save(comment).Error
}
