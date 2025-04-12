package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
)
//faudrait peut-être créer la session ici et après dans login,register,et googleregister transferer la session

// /backend/handler/guest.go
// execute le template du home pour guest
func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités b")
	userID := GetUserIDFromSession(r)

	var isAuthenticated bool
	if userID != 0 {
		isAuthenticated = true
	}

	var posts []structs.Post

	//j'ai encore un peu de mal à comprendre les preload, donc faudra que je demande tout de même pcq même en lisant je capte que dalle

	err := db.DB.Preload("Comment").Preload("Comment.UserID").Preload("Comment.CommentsLike").Find(&posts).Error
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

	var users []structs.User
	err = db.DB.Find(&users).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des utilisateurs:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}
	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", struct {
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
// /backend/handler/guest.go
// execute le template de la catégorie "hack"
func GuestHack(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPostsByCategory("hack")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts Hack", http.StatusInternalServerError)
		return
	}
	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", posts)
}
// /backend/handler/guest.go
// execute le template de la catégorie "prog"
func GuestProg(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPostsByCategory("prog")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts Prog", http.StatusInternalServerError)
		return
	}
	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", posts)
}
// /backend/handler/guest.go
// execute le template de la catégorie "news"
func GuestNews(w http.ResponseWriter, r *http.Request) {
	posts, err := GetPostsByCategory("news")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des posts News", http.StatusInternalServerError)
		return
	}
	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", posts)
}

// /backend/handler/guest.go
// cherche des postes comprenant dans leur commentaire le mot clé
func GuestSearch(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	posts, err := SearchPosts(searchQuery)
	if err != nil {
		http.Error(w, "Erreur lors de la recherche de posts", http.StatusInternalServerError)
		return
	}

	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", posts)
}