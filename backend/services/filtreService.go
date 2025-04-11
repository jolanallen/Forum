package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"log"
	"net/http"
	"strconv"
)

func FilterPostsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
			return
		}
		categoryIDStr := r.FormValue("categoryID")
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
			return
		}
		var posts []structs.Post
		err = db.DB.
			Preload("Comments").
			Preload("Comments.User").
			Preload("Comments.CommentsLike").
			Preload("Category").
			Preload("Image").
			Find(&posts, "category_id = ?", categoryID).Error
		if err != nil {
			log.Println("Erreur lors de la récupération des posts:", err)
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
			return
		}

		userID := GetUserIDFromSession(r)
		isAuthenticated := userID != 0

		var categories []structs.Category
		err = db.DB.Find(&categories).Error
		if err != nil {
			log.Println("Erreur lors de la récupération des catégories:", err)
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
			return
		}

		var users []structs.User
		err = db.DB.Find(&users).Error
		if err != nil {
			log.Println("Erreur lors de la récupération des utilisateurs:", err)
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
			return
		}

		Templates.ExecuteTemplate(w, "home_guest.html", struct {
			IsAuthenticated bool
			Posts           []structs.Post
			Categories      []structs.Category
			Users           []structs.User
		}{
			IsAuthenticated: isAuthenticated,
			Posts:           posts,
			Categories:      categories,
			Users:           users,
		})
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
}
