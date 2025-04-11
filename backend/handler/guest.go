package handler

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités b")
	userID := utils.GetUserIDFromSession(r)

	// Vérifier si l'utilisateur est authentifié
	var isAuthenticated bool
	if userID != 0 {
		isAuthenticated = true
	}

	// Récupérer les posts, commentaires et likes (comme précédemment)
	var posts []structs.Post
	err := db.DB.Preload("Comments").Preload("Comments.User").Preload("Comments.CommentsLike").Preload("Comments.CommentsDislike").Find(&posts).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des posts:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	// Récupérer toutes les catégories
	var categories []structs.Category
	err = db.DB.Find(&categories).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des catégories:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	// Créer un template HTML et passer les données nécessaires
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Erreur lors du parsing du template:", err)
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}

	// Afficher le template avec les données des posts et catégories
	err = tmpl.Execute(w, struct {
		Posts           []structs.Post
		Categories      []structs.Category
		IsAuthenticated bool
	}{Posts: posts, Categories: categories, IsAuthenticated: isAuthenticated})
	if err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}

func GuestHack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Hack pour les invités")
}

func GuestProg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Programme pour les invités")
}

func GuestNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page des actualités pour les invités")
}

func GuestSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page de recherche pour les invités")
}

func GuestAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page À propos pour les invités")
}
