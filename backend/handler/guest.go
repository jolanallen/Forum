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


func GuestHome(w http.ResponseWriter, r *http.Request) {
	userID := services.GetUserIDFromSession(r)
	isAuthenticated := userID != 0

	// Récupérer les posts avec les commentaires, leurs auteurs et leurs likes
	var posts []structs.Post
	if err := db.DB.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User"). // <-- Correct ici si struct Comment a une relation User
		Find(&posts).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts")
		return
	}


	// Récupération des catégories
	var categories []structs.Category
	if err := db.DB.Find(&categories).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des catégories")
		return
	}

	// Récupérer l'utilisateur connecté si authentifié
	var user structs.User
	if isAuthenticated {
		if err := db.DB.First(&user, userID).Error; err != nil {
			handleError(w, err, "Erreur lors de la récupération de l'utilisateur")
			return
		}
	}

	// Récupération des autres utilisateurs (si vraiment nécessaire)
	var users []structs.User
	if err := db.DB.Find(&users).Error; err != nil {
		handleError(w, err, "Erreur lors de la récupération des utilisateurs")
		return
	}

	// Ajouter le champ ActivePage ici pour indiquer la page active
	services.RenderTemplate(w, "forum/home.html", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
		Users           []structs.User
		User            structs.User
		ActivePage      string  // <-- Ajouté ici
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
		Users:           users,
		User:            user,
		ActivePage:      "home",  // <-- Page active définie ici
	})
}




// Catégorie Hack
func CategoryHack(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("hack")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Hack")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_hack.html", posts)
}

// Catégorie Programmation
func CategoryProg(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("prog")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Prog")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_prog.html", posts)
}

// Catégorie News
func CategoryNews(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("news")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts News")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_news.html", posts)
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
