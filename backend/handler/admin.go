package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	if adminID == nil {
		http.Error(w, "Non autorisé", http.StatusUnauthorized)
		return
	}

	adminIDUint, ok := adminID.(uint64)
	if !ok {
		http.Error(w, "Erreur ID admin", http.StatusBadRequest)
		return
	}

	data, err := services.GetAdminDashboardData(adminIDUint)
	if err != nil {
		http.Error(w, "Erreur lors du chargement du dashboard", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin_dashboard.html")
	if err != nil {
		http.Error(w, "Erreur template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("userID")
	var admin structs.Admin

	// Récupérer l'admin à partir de la base de données
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Accès interdit", http.StatusForbidden)
		} else {
			http.Error(w, "Erreur d'accès", http.StatusInternalServerError)
		}
		return
	}

	// Récupérer les paramètres de l'URL
	vars := mux.Vars(r)
	userIDStr := vars["id"]

	// Convertir l'ID de l'utilisateur
	userIDToDelete, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID utilisateur invalide", http.StatusBadRequest)
		return
	}

	// Vérifier si l'utilisateur existe
	var user structs.User
	row = db.DB.QueryRow("SELECT userID, userEmail FROM users WHERE userID = ?", userIDToDelete)
	if err := row.Scan(&user.UserID, &user.UserEmail); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		} else {
			http.Error(w, "Erreur lors de la recherche de l'utilisateur", http.StatusInternalServerError)
		}
		return
	}

	// Supprimer l'utilisateur de la base de données
	_, err = db.DB.Exec("DELETE FROM users WHERE userID = ?", userIDToDelete)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression de l'utilisateur", http.StatusInternalServerError)
		return
	}

	// Rediriger vers le dashboard
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Accès interdit", http.StatusForbidden)
		} else {
			http.Error(w, "Erreur d'accès", http.StatusInternalServerError)
		}
		return
	}

	// Récupérer les paramètres de l'URL
	vars := mux.Vars(r)
	commentIDStr := vars["id"]

	// Convertir l'ID du commentaire
	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID commentaire invalide", http.StatusBadRequest)
		return
	}

	// Vérifier si le commentaire existe
	var comment structs.Comment
	row = db.DB.QueryRow("SELECT commentID, content FROM comments WHERE commentID = ?", commentID)
	if err := row.Scan(&comment.CommentID, &comment.CommentContent); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		} else {
			http.Error(w, "Erreur lors de la recherche du commentaire", http.StatusInternalServerError)
		}
		return
	}

	// Supprimer le commentaire de la base de données
	_, err = db.DB.Exec("DELETE FROM comments WHERE commentID = ?", commentID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du commentaire", http.StatusInternalServerError)
		return
	}

	// Rediriger vers le dashboard
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin

	// Vérifier l'existence de l'admin
	row := db.DB.QueryRow("SELECT adminID FROM admins WHERE adminID = ?", adminID)
	if err := row.Scan(&admin.AdminID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Accès interdit", http.StatusForbidden)
		} else {
			http.Error(w, "Erreur d'accès", http.StatusInternalServerError)
		}
		return
	}

	// Récupérer les paramètres de l'URL
	vars := mux.Vars(r)
	postIDStr := vars["id"]

	// Convertir l'ID du post
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID post invalide", http.StatusBadRequest)
		return
	}

	// Vérifier si le post existe
	var post structs.Post
	row = db.DB.QueryRow("SELECT postID FROM posts WHERE postID = ?", postID)
	if err := row.Scan(&post.PostID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post introuvable", http.StatusNotFound)
		} else {
			http.Error(w, "Erreur lors de la recherche du post", http.StatusInternalServerError)
		}
		return
	}

	// Supprimer le post
	_, err = db.DB.Exec("DELETE FROM posts WHERE postID = ?", postID)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du post", http.StatusInternalServerError)
		return
	}

	// Rediriger
	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
