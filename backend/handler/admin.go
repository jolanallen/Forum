package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
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

	if err := db.DB.Where("adminID = ?", adminID).First(&admin).Error; err != nil {
		http.Error(w, "Accès interdit", http.StatusForbidden)
		return
	}
	vars := mux.Vars(r)
	userIDStr := vars["id"]

	userIDToDelete, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID utilisateur invalide", http.StatusBadRequest)
		return
	}

	var user structs.User
	if err := db.DB.Where("userID = ?", userIDToDelete).First(&user).Error; err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression de l'utilisateur", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin
	if err := db.DB.Where("adminID = ?", adminID).First(&admin).Error; err != nil {
		http.Error(w, "Accès interdit", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	commentIDStr := vars["id"]

	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID commentaire invalide", http.StatusBadRequest)
		return
	}

	var comment structs.Comment
	if err := db.DB.Where("commentID = ?", commentID).First(&comment).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression du commentaire", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}

func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	adminID := r.Context().Value("adminID")
	var admin structs.Admin
	if err := db.DB.Where("adminID = ?", adminID).First(&admin).Error; err != nil {
		http.Error(w, "Accès interdit", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	postIDStr := vars["id"]

	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID post invalide", http.StatusBadRequest)
		return
	}

	var post structs.Post
	if err := db.DB.Where("postID = ?", postID).First(&post).Error; err != nil {
		http.Error(w, "Post introuvable", http.StatusNotFound)
		return
	}

	if err := db.DB.Delete(&post).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression du post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
