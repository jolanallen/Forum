package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"net/http"
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
