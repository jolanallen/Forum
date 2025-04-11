package utils

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func GetUserIDFromUsername(username string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("users_username = ?", username).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromEmail(email string) (uint64, error) {
	var user structs.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return 0, fmt.Errorf("utilisateur introuvable")
	}
	return user.UserID, nil
}
func GetUserIDFromCommentID(commentID uint64) (uint64, error) {
	var comment structs.Comment
	if err := db.DB.Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		return 0, fmt.Errorf("commentaire introuvable")
	}
	return comment.UserID, nil
}
func GetUserIDFromPostID(postID uint64) (uint64, error) {
	var post structs.Post
	if err := db.DB.Where("post_id = ?", postID).First(&post).Error; err != nil {
		return 0, fmt.Errorf("post introuvable")
	}
	return post.UserID, nil
}

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

		userID := utils.GetUserIDFromSession(r)
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

		utils.Templates.ExecuteTemplate(w, "home_guest.html", struct {
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

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur recherché depuis les paramètres de la requête
	username := r.URL.Query().Get("username")

	var users []structs.User
	if err := db.DB.Where("username LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
		//erreur 500
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	//FAUT FAIRE EN SORTE QU4IL RENVOIE VERS LE TEMPLATE DU PROFIL DE L4USER CHERCHER

	//OU ALORS ON LES PARSE DIRECT

	// Créer un template HTML et afficher les résultats
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}

	// Passer les utilisateurs trouvés au template
	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}
