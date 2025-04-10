package backend

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
)

// pour les COOKIES
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.Cookie c'est les informations directe récuperer par le navigateur
		cookie, err := r.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		var session structs.Session
		result := db.DB.Where("session_token = ?", cookie.Value).First(&session)
		if result.Error != nil || session.ExpiresAt.Before(time.Now()) {
			http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", session.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// POUR LES ADMIN
// on verifie que l'userID est ou n'est pas un admin
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("userID")
		var admin structs.Admin
		if err := db.DB.Where("user_id = ?", userID).First(&admin).Error; err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
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

// fonction pour ajouter un like au post
func UserLikePost(w http.ResponseWriter, r *http.Request) {
	//on récup dans la requete uniquement l'id du post et si on le trouve on l'incrémente de 1
	postIDStr := strings.TrimPrefix(r.URL.Path, "/user/post/")
	postIDStr = strings.TrimSuffix(postIDStr, "/like")
	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var post structs.Post
	if err := db.DB.First(&post, postID).Error; err != nil {
		http.Error(w, "Post introuvable", http.StatusNotFound)
		return
	}

	post.PostLike++
	//on le met à jour
	db.DB.Save(&post)
	//redirection vers home mais a voir pcq c'est surement pas ça
	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

func UserAddComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		//pareil que pour post
		postIDStr := strings.TrimPrefix(r.URL.Path, "/user/post/")
		postIDStr = strings.TrimSuffix(postIDStr, "/comment")
		postID, err := strconv.ParseUint(postIDStr, 10, 64)
		if err != nil {
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}

		userID := r.Context().Value("userID").(uint64)
		content := r.FormValue("comment")

		comment := structs.Comment{
			UserID:  userID,
			PostID:  postID,
			Content: content,
			Status:  "published",
			Visible: true,
		}

		if err := db.DB.Create(&comment).Error; err != nil {
			http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forum/", http.StatusSeeOther)
	} else {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
}

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		userID := r.Context().Value("userID").(uint64)
		postKey := uuid.New().String()
		content := r.FormValue("content")

		// Récupération du fichier image
		file, header, err := r.FormFile("image")
		var imageID *uint64
		if err == nil {
			defer file.Close()
			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, file); err == nil {
				image := structs.Image{
					Filename: header.Filename,
					Data:     buf.Bytes(),
					URL:      "/images/" + header.Filename,
				}
				if err := db.DB.Create(&image).Error; err == nil {
					imageID = &image.ImageID
				}
			}
		}

		post := structs.Post{
			PostKey:     postKey,
			PostComment: content,
			UserID:      userID,
			ImageID:     imageID,
		}

		if err := db.DB.Create(&post).Error; err != nil {
			//erreur 500
			http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
			return
		}
		//tjr pas convaincu de la route
		http.Redirect(w, r, "/forum/", http.StatusSeeOther)
	} else {
		// et si c'est pas une requete post on le vire
		utils.Templates.ExecuteTemplate(w, "create_post.html", nil)
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

// Récupérer tous les users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(users)
}

// Récupérer tous les admins
// ça ne return rien, c'est juste pour la prog
func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
}

// CreateSession crée une nouvelle session pour l'utilisateur
func CreateSession(userID uint64) (string, error) {
	sessionToken := utils.GenerateToken()
	expiration := time.Now().Add(24 * time.Hour)

	// Créer une nouvelle session
	session := structs.Session{
		UserID:       userID,
		SessionToken: sessionToken,
		ExpiresAt:    expiration,
	}

	// Insérer la session dans la base de données
	result := db.DB.Create(&session)
	if result.Error != nil {
		return "", result.Error
	}

	return sessionToken, nil
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
		sessionToken, err := CreateSession(user.UserID)
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

//le pb c'est qu'il prend par rapport à la session de son nav et pas de notre bdd
func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0 // Pas de session
	}
	var userID uint64
	fmt.Sscanf(cookie.Value, "%d", &userID)
	return userID
}

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités b")
	userID := GetUserIDFromSession(r)

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

func FilterPostsByCategory(w http.ResponseWriter, r *http.Request) {
	categoryIDStr := strings.TrimPrefix(r.URL.Path, "/category/")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de catégorie invalide", http.StatusBadRequest)
		return
	}

	var posts []structs.Post
	err = db.DB.Preload("Comments").Preload("Comments.User").Preload("Category").Where("category_id = ?", categoryID).Find(&posts).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des posts:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	userID := GetUserIDFromSession(r)
	var isAuthenticated bool
	if userID != 0 {
		isAuthenticated = true
	}

	var categories []structs.Category
	err = db.DB.Find(&categories).Error
	if err != nil {
		log.Println("Erreur lors de la récupération des catégories:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	utils.Templates.ExecuteTemplate(w, "home_guest.html", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
	})
}
