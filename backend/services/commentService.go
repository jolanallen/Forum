package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"database/sql"
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
	// Extraction de l'ID du post depuis l'URL
	postID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Récupération de l'ID utilisateur depuis le contexte
	userID := r.Context().Value("userID").(uint64)
	// Récupération du contenu du commentaire
	content := r.FormValue("comment")

	// Création du commentaire
	comment := structs.Comment{
		UserID:  userID,
		PostID:  postID,
		CommentContent: content,
		CommentStatus:  "published",
		CommentVisible: true,
	}

	// Insertion du commentaire dans la base de données avec une requête SQL
	query := "INSERT INTO comments (userID, postID, content, status, visible) VALUES (?, ?, ?, ?, ?)"
	_, err = db.DB.Exec(query, comment.UserID, comment.PostID, comment.CommentContent, comment.CommentStatus, comment.CommentVisible)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
		return
	}

	// Redirection après l'ajout du commentaire
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func UserEditComment(w http.ResponseWriter, r *http.Request) {
	// Extraction de l'ID du commentaire depuis l'URL
	commentID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	// Recherche du commentaire dans la base de données
	var comment structs.Comment
	query := "SELECT commentID, userID, postID, content, status, visible FROM comments WHERE commentID = ?"
	err = db.DB.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.CommentContent, &comment.CommentStatus, &comment.CommentVisible)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Commentaire introuvable", http.StatusNotFound)
			return
		}
		http.Error(w, "Erreur de récupération du commentaire", http.StatusInternalServerError)
		return
	}

	// Vérification de l'autorisation de l'utilisateur
	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	// Récupération du nouveau contenu du commentaire
	updatedContent := r.FormValue("comment")
	comment.CommentContent = updatedContent

	// Mise à jour du commentaire dans la base de données
	updateQuery := "UPDATE comments SET content = ? WHERE commentID = ?"
	_, err = db.DB.Exec(updateQuery, comment.CommentContent, comment.CommentID)
	if err != nil {
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}

	// Redirection après la mise à jour
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func UserDeleteComment(w http.ResponseWriter, r *http.Request) {
	// Extraction de l'ID du commentaire depuis l'URL
	commentID, err := ExtractIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	// Recherche du commentaire dans la base de données
	var comment structs.Comment
	query := "SELECT commentID, userID, postID, content, status, visible FROM comments WHERE commentID = ?"
	err = db.DB.QueryRow(query, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.CommentContent, &comment.CommentStatus, &comment.CommentVisible)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Commentaire introuvable", http.StatusNotFound)
			return
		}
		http.Error(w, "Erreur de récupération du commentaire", http.StatusInternalServerError)
		return
	}

	// Vérification de l'autorisation de l'utilisateur
	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	// Suppression du commentaire dans la base de données
	deleteQuery := "DELETE FROM comments WHERE commentID = ?"
	_, err = db.DB.Exec(deleteQuery, comment.CommentID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	// Redirection après la suppression
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}
