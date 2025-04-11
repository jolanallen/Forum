package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"net/http"

	"gorm.io/gorm"
)

func ToggleLikePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	postID, err := ExtractPostIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	post, err := getPostByID(postID)
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
		if err := removeLike(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := addLike(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

func hasUserLikedPost(userID, postID uint64) (bool, error) {
	var like structs.PostLike
	err := db.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func addLike(userID, postID uint64, post *structs.Post) error {
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

func removeLike(userID, postID uint64, post *structs.Post) error {
	if err := db.DB.Where("userID = ? AND postID = ?", userID, postID).Delete(&structs.PostLike{}).Error; err != nil {
		return err
	}
	if post.PostLike > 0 {
		post.PostLike--
	}
	return db.DB.Save(post).Error
}
