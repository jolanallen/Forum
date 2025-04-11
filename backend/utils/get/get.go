package get

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités b")
	userID := utils.GetUserIDFromSession(r)

	// Vérifier si l'utilisateur est authentifié
	var isAuthenticated bool
	if userID != 0 {
		isAuthenticated = true
	}

	var posts []structs.Post
	err := db.DB.Preload("Comment").Preload("Comment.UserID").Preload("Comment.CommentsLike").Find(&posts).Error
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

	// Récupérer toutes les images et utilisateurs
	var users []structs.User
	err = db.DB.Find(&users).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des utilisateurs:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}
	// On passe tous les résultats au template
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
}

// lors du login ( changer peut-etre le nom de la fonction en fonction de la route)
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// récupération formulaire
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user structs.User
		//vérifier pour le .ERROR
		if err := db.DB.Where("users_username = ?", username).First(&user).Error; err != nil {
			//erreur 401
			http.Error(w, "Utilisateur inconnu", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(password, user.UserPasswordHash) {
			//erreur 401 si pwd par bon
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}
		//crée un token de session pour l'utilisateur
		sessionToken, err := utils.CreateSession(user.UserID)
		if err != nil {
			http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
			return
		}

		//on l'insert dans le nav du client (cookie)
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})
		//redirection vers home et http.StatusSeeOther sert au cas où il y aura rafraichissement de la page ( status de réussite code 303)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		// Si la méthode n'est pas POST, on affiche le formulaire de connexion
		utils.Templates.ExecuteTemplate(w, "login.html", nil)
	}
}
// route, il faudra que je vérifie si c'est le bon nom et renvoie vers le bon template () aussi,
// fonction qui renvoie tout les posts de la base de données
func GuestHomeHandler(w http.ResponseWriter, r *http.Request) {
	var posts []structs.Post
	db.DB.Preload("User").Preload("Comments").Find(&posts)
	// faut changer vers le bon template
	utils.Templates.ExecuteTemplate(w, "home.html", posts)
}
