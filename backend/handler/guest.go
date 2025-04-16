package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, err error, message string) {
	log.Println(message, err)
	http.Error(w, message, http.StatusInternalServerError)
}

func GuestHome(w http.ResponseWriter, r *http.Request) {
	userID := services.GetUserIDFromSession(r)
	isAuthenticated := userID != 0
	fmt.Println(services.HashPassword("hashed_password_1"))

	// Récupérer les posts avec les commentaires, leurs auteurs et leurs likes (SQL pur)
	var posts []structs.Post
	rows, err := db.DB.Query(`
		SELECT posts.postID, posts.categoryID, posts.postKey, posts.imageID, posts.postComment, posts.postLike, posts.createdAt, 
		       users.userUsername
		FROM posts
		JOIN users ON posts.userID = users.userID
	`)
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var username string // Variable temporaire pour stocker le nom d'utilisateur
		// Récupérer toutes les colonnes et inclure le nom d'utilisateur
		if err := rows.Scan(&post.PostID, &post.CategoryID, &post.PostKey, &post.ImageID, &post.PostComment, &post.PostLike, &post.PostCreatedAt, &username); err != nil {
			handleError(w, err, "Erreur lors de la récupération des données de post")
			return
		}
		// Associer le nom d'utilisateur récupéré à un champ d'un autre objet, si nécessaire
		post.UserUsername = username // Pas besoin de l'ajouter à la structure, juste l'utiliser localement
		posts = append(posts, post)
	}

	// Récupérer les catégories (SQL pur)
	var categories []structs.Category
	rows, err = db.DB.Query("SELECT categoryID, categoryName FROM categories")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des catégories")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			handleError(w, err, "Erreur lors de la récupération des catégories")
			return
		}
		categories = append(categories, category)
	}

	// Récupérer l'utilisateur connecté si authentifié (SQL pur)
	var user structs.User
	if isAuthenticated {
		row := db.DB.QueryRow("SELECT userID, userUsername FROM users WHERE userID = $1", userID)
		if err := row.Scan(&user.UserID, &user.UserUsername); err != nil {
			if err == sql.ErrNoRows {
				handleError(w, err, "Utilisateur non trouvé")
			} else {
				handleError(w, err, "Erreur lors de la récupération de l'utilisateur")
			}
			return
		}
	}

	// Ajouter le champ ActivePage ici pour indiquer la page active
	services.RenderTemplate(w, "forum/home.html", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
		User            structs.User
		ActivePage      string
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
		User:            user,
		ActivePage:      "home",
	})
}

func CategoryHack(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("hack")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Hack")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_hack.html", posts)
}

func CategoryProg(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("prog")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts Prog")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_prog.html", posts)
}

func CategoryNews(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("news")
	if err != nil {
		handleError(w, err, "Erreur lors de la récupération des posts News")
		return
	}
	services.RenderTemplate(w, "guest/catégorie_news.html", posts)
}
func SearchPseudo(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	posts, err := services.SearchPosts(searchQuery)
	if err != nil {
		handleError(w, err, "Erreur lors de la recherche de posts")
		return
	}
	services.RenderTemplate(w, "guest/search.html", posts)
}
func AboutForum(w http.ResponseWriter, r *http.Request) {
	services.RenderTemplate(w, "guest/about.html", nil)
}
