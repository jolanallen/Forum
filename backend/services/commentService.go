package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"net/http"
)

func HandleCommentActions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		UserAddComment(w, r)
	case http.MethodPut:
		UserEditComment(w, r)
	case http.MethodDelete:
		UserDeleteComment(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
func UserAddComment(w http.ResponseWriter, r *http.Request) {
	postID, err := ExtractPostIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	content := r.FormValue("comment")

	comment := structs.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
		Status:  "published",
		Visible: true,
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func UserEditComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := ExtractCommentIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	var comment structs.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	updatedContent := r.FormValue("comment")
	comment.Content = updatedContent

	if err := db.DB.Save(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}




func UserDeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := ExtractCommentIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	var comment structs.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	if err := db.DB.Delete(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}
