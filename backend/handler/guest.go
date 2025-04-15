package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"log"
	"net/http"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	userID := services.GetUserIDFromSession(r)

	var isAuthenticated bool
	if userID != 0 {
		isAuthenticated = true
	}

	var posts []structs.Post

	err := db.DB.Preload("Comments").Preload("Comments.UserID").Preload("Comments.CommentsLike").Find(&posts).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des posts:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	var categories []structs.Category
	err = db.DB.Find(&categories).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des catégories:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	user := structs.User{}
	if isAuthenticated {
		err = db.DB.First(&user, userID).Error
		if err != nil {
			log.Println("Erreur lors de la récupération de l'utilisateur:", err)
			http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
			return
		}
	}
	var users []structs.User
	err = db.DB.Find(&users).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des utilisateurs:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}
	log.Println("IsAuthenticated:", isAuthenticated)
	log.Println("Posts:", posts)
	log.Println("Categories:", categories)
	log.Println("Users:", users)
	log.Println("User:", user)
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

func CategoryHack(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("hack")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts Hack", http.StatusInternalServerError)
		return
	}
	services.RenderTemplate(w, "guest/category_hack.html", posts)
}

func CategoryProg(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("prog")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts Prog", http.StatusInternalServerError)
		return
	}
	services.RenderTemplate(w, "guest/category_prog.html", posts)
}

func CategoryNews(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("news")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts News", http.StatusInternalServerError)
		return
	}
	services.RenderTemplate(w, "guest/category_news.html", posts)
}

func SearchPseudo(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	posts, err := services.SearchPosts(searchQuery)
	if err != nil {
		http.Error(w, "Erreur lors de la recherche de posts", http.StatusInternalServerError)
		return
	}

	services.RenderTemplate(w, "guest/search.html", posts)
}

func AboutForum(w http.ResponseWriter, r *http.Request) {
	services.RenderTemplate(w, "guest/about.html", nil)
}
 
 import (
    "Forum/backend/db"
    "Forum/backend/services"
    "Forum/backend/structs"
    "log"
    "net/http"
 )
 
 func GuestHome(w http.ResponseWriter, r *http.Request) {
    userID := services.GetUserIDFromSession(r)
 
    var isAuthenticated bool
    if userID != 0 {
        isAuthenticated = true
    }
 
    var posts []structs.Post
 
    err := db.DB.Preload("Comments").Preload("Comments.UserID").Preload("Comments.CommentsLike").Find(&posts).Error
    if err != nil {
        log.Println("Erreur lors de la récupération des posts:", err)
        http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
        return
    }
 
    var categories []structs.Category
    err = db.DB.Find(&categories).Error
    if err != nil {
        log.Println("Erreur lors de la récupération des catégories:", err)
        http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
        return
    }
 
    user := structs.User{}
    if isAuthenticated {
        err = db.DB.First(&user, userID).Error
        if err != nil {
            log.Println("Erreur lors de la récupération de l'utilisateur:", err)
            http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
            return
        }
    }
    var users []structs.User
    err = db.DB.Find(&users).Error
    if err != nil {
        log.Println("Erreur lors de la récupération des utilisateurs:", err)
        http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
        return
    }
    log.Println("IsAuthenticated:", isAuthenticated)
    log.Println("Posts:", posts)
    log.Println("Categories:", categories)
    log.Println("Users:", users)
    log.Println("User:", user)
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
 
 func CategoryHack(w http.ResponseWriter, r *http.Request) {
    posts, err := services.GetPostsByCategory("hack")
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des posts Hack", http.StatusInternalServerError)
        return
    }
    services.RenderTemplate(w, "guest/category_hack.html", posts)
 }
 
 func CategoryProg(w http.ResponseWriter, r *http.Request) {
    posts, err := services.GetPostsByCategory("prog")
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des posts Prog", http.StatusInternalServerError)
        return
    }
    services.RenderTemplate(w, "guest/category_prog.html", posts)
 }
 
 func CategoryNews(w http.ResponseWriter, r *http.Request) {
    posts, err := services.GetPostsByCategory("news")
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des posts News", http.StatusInternalServerError)
        return
    }
    services.RenderTemplate(w, "guest/category_news.html", posts)
 }
 
 func SearchPseudo(w http.ResponseWriter, r *http.Request) {
    searchQuery := r.URL.Query().Get("query")
    posts, err := services.SearchPosts(searchQuery)
    if err != nil {
        http.Error(w, "Erreur lors de la recherche de posts", http.StatusInternalServerError)
        return
    }
 
    services.RenderTemplate(w, "guest/search.html", posts)
 }
 
 func AboutForum(w http.ResponseWriter, r *http.Request) {
    services.RenderTemplate(w, "guest/about.html", nil)
 }
