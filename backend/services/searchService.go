package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
			return
		}
		username := r.FormValue("userUsername")
		var users []structs.User
		if err := db.DB.Where("userUsername LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
			return
		}

		if len(users) == 1 {
			http.Redirect(w, r, "/profile/"+users[0].UserUsername, http.StatusSeeOther)
			return
		}
		tmpl, err := template.ParseFiles("templates/search_results.html")
		if err != nil {
			http.Error(w, "Erreur de template", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, users)
		if err != nil {
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
}

func SearchPosts(query string) ([]structs.Post, error) {
	var posts []structs.Post
	err := db.DB.Where("postComment LIKE ?", "%"+query+"%").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func FilterPostsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
			return
		}
		categoryIDStr := r.FormValue("categoriesID")
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
			Find(&posts, "categoriesID = ?", categoryID).Error
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

		RenderTemplate(w, "BoyWithUke_Prairies", struct {
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
