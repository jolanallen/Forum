package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"log"
	"net/http"
)

// Fonction utilitaire pour gérer les erreurs
func handleError(w http.ResponseWriter, err error, message string) {
	log.Println(message, err)
	http.Error(w, message, http.StatusInternalServerError)
}

// Fonction d'affichage de la page d'accueil des invités
func GuestHome(w http.ResponseWriter, r *http.Request) {
	userID := services.GetUserIDFromSession(r)
	isAuthenticated := userID != 0

	var posts []structs.Post
	if err := db.DB.Preload("Comments").Preload("Comments.UserID").Preload("Comments.CommentsLike").Find(&posts).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts")
		return
	}

	var categories []structs.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des catégories")
		return
	}

	var user structs.User
	if isAuthenticated {
		if err := db.DB.First(&user, userID).Error; err != nil {
			handleError(w, err, "Erreur lors de la récupération de l'utilisateur")
			return
		}
	}

	// Ne récupère pas tous les utilisateurs si ce n'est pas nécessaire
	var users []structs.User
	if err := db.DB.Find(&users).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des utilisateurs")
		return
	}

	// Passer toutes les données nécessaires au template
	services.RenderTemplate(w, "forum/home.html", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
		Users           []structs.User
		User            structs.User
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
		Users:           users,
		User:            user,
	})
}

// Catégorie Hack
func CategoryHack(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("hack")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Hack")
		return
	}
	services.RenderTemplate(w, "guest/category_hack.html", posts)
}

// Catégorie Programmation
func CategoryProg(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("prog")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Prog")
		return
	}
	services.RenderTemplate(w, "guest/category_prog.html", posts)
}

// Catégorie News
func CategoryNews(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("news")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts News")
		return
	}
	services.RenderTemplate(w, "guest/category_news.html", posts)
}

// Recherche de posts par pseudo
func SearchPseudo(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	posts, err := services.SearchPosts(searchQuery)
	if err != nil {
		handleError(w, err, "Erreur lors de la recherche de posts")
		return
	}

	services.RenderTemplate(w, "guest/search.html", posts)
}

// Page "À propos" du forum
func AboutForum(w http.ResponseWriter, r *http.Request) {
	services.RenderTemplate(w, "guest/about.html", nil)
}
