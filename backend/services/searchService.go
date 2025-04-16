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
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	username := r.FormValue("userUsername")
	rows, err := db.DB.Query("SELECT userID, userUsername FROM users WHERE userUsername LIKE ?", "%"+username+"%")
	if err != nil {
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		log.Println("Query erreur:", err)
		return
	}
	defer rows.Close()

	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.UserID, &user.UserUsername); err != nil {
			log.Println("Erreur scan:", err)
			continue
		}
		users = append(users, user)
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
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}

func SearchPosts(query string) ([]structs.Post, error) {
	rows, err := db.DB.Query("SELECT postID, postComment, userID, postKey, imageID, postLike, categoryID, createdAt FROM posts WHERE postComment LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []structs.Post
	for rows.Next() {
		var p structs.Post
		err := rows.Scan(&p.PostID, &p.PostComment, &p.UserID, &p.PostKey, &p.ImageID, &p.PostLike, &p.CategoryID, &p.PostCreatedAt)
		if err != nil {
			log.Println("Erreur scan post:", err)
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func FilterPostsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	categoryIDStr := r.FormValue("categoriesID")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
		return
	}

	postRows, err := db.DB.Query(`SELECT postID, postComment, userID, postKey, imageID, postLike, categoryID, createdAt FROM posts WHERE categoryID = ?`, categoryID)
	if err != nil {
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		log.Println("Erreur récupération des posts:", err)
		return
	}
	defer postRows.Close()

	var posts []structs.Post
	for postRows.Next() {
		var post structs.Post
		if err := postRows.Scan(&post.PostID, &post.PostComment, &post.UserID, &post.PostKey, &post.ImageID, &post.PostLike, &post.CategoryID, &post.PostCreatedAt); err != nil {
			log.Println("Erreur scan post:", err)
			continue
		}
		posts = append(posts, post)
	}

	userID := GetUserIDFromSession(r)
	isAuthenticated := userID != 0

	// Récupérer les catégories
	catRows, err := db.DB.Query("SELECT categoryID, categoryName, categoryDescription FROM categories")
	if err != nil {
		http.Error(w, "Erreur récupération des catégories", http.StatusInternalServerError)
		return
	}
	defer catRows.Close()

	var categories []structs.Category
	for catRows.Next() {
		var c structs.Category
		if err := catRows.Scan(&c.CategoryID, &c.CategoryName, &c.CategoryDescription); err != nil {
			continue
		}
		categories = append(categories, c)
	}

	// Récupérer les utilisateurs
	userRows, err := db.DB.Query("SELECT userID, userUsername FROM users")
	if err != nil {
		http.Error(w, "Erreur récupération des utilisateurs", http.StatusInternalServerError)
		return
	}
	defer userRows.Close()

	var users []structs.User
	for userRows.Next() {
		var u structs.User
		if err := userRows.Scan(&u.UserID, &u.UserUsername); err != nil {
			continue
		}
		users = append(users, u)
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
}
