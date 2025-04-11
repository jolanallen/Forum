package backend
/*
import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
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

// lors du login ( changer peut-etre le nom de la fonction en fonction de la route)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
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

		if !utils.CheckPasswordHash(password, user.PasswordHash) {
			//erreur 401 si pwd par bon
			http.Error(w, "Mot de passe invalide", http.StatusUnauthorized)
			return
		}
		//crée un token de session pour l'utilisateur
		token := utils.GenerateToken()
		// Crée une nouvelle session dans la base de données
		//SessionsID et CreatAt sont automatiques
		session := structs.Session{
			UserID:       user.UsersID,
			SessionToken: token,
			ExpiresAt:    time.Now().Add(24 * time.Hour),
		}
		//on l'insert dans la base de données (commande gorm)
		db.DB.Create(&session)
		//on l'insert dans le nav du client (cookie)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   token,
			Expires: session.ExpiresAt,
		})
		//redirection vers home et http.StatusSeeOther sert au cas où il y aura rafraichissement de la page ( status de réussite code 303)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		// Si la méthode n'est pas POST, on affiche le formulaire de connexion
		utils.Templates.ExecuteTemplate(w, "login.html", nil)
	}
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

	post.PostLikes++
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

		userID := r.Context().Value("user_id").(uint64)
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
	if r.Method != http.MethodPost {
		utils.Templates.ExecuteTemplate(w, "create_post.html", nil)
		return
	}

	userID := r.Context().Value("user_id").(uint64)
	postKey := uuid.New().String()
	content := r.FormValue("content")

	// Récupération du fichier image
	file, header, err := r.FormFile("image")
	var imageID *uint
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
		PostKey:      postKey,
		PostComments: content,
		UserID:       userID,
		ImageID:      imageID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur recherché depuis les paramètres de la requête
	username := r.URL.Query().Get("username")

	// Chercher l'utilisateur dans la base de données
	var users []structs.User
	err := db.DB.Where("username LIKE ?", "%"+username+"%").Find(&users).Error
	if err != nil {
		log.Println("Erreur lors de la recherche des utilisateurs:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	// Créer un template HTML et afficher les résultats
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		log.Println("Erreur lors du parsing du template:", err)
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}

	// Passer les utilisateurs trouvés au template
	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []structs.User
	result := db.DB.Find(&users) // Récupérer tous les users
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(users)
}
func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins) // Récupérer tous les admins
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
}

// generateSessionToken génère un token de session aléatoire
func generateSessionToken() string {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(token)
}

// CreateSession crée une nouvelle session pour l'utilisateur
func CreateSession(userID uint) (string, error) {
	sessionToken := generateSessionToken()
	expiration := time.Now().Add(24 * time.Hour)

	// Créer une nouvelle session
	session := structs.Session{
		UserID:       uint64(userID),
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

// Login gère la logique de connexion de l'utilisateur
func Login(w http.ResponseWriter, r *http.Request) {
	// Ici, vous devez vérifier les identifiants de l'utilisateur et récupérer son userID
	// Supposons que vous avez récupéré le userID de l'utilisateur après la validation des identifiants

	var userID uint // Remplacez cette ligne par la récupération du userID réel
	// Exemple: userID = 1

	// Créer une session pour l'utilisateur
	sessionToken, err := CreateSession(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la création de la session", http.StatusInternalServerError)
		return
	}

	// Créer un cookie avec le token de session
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// Rediriger l'utilisateur vers la page d'accueil
	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

// Register gère l'inscription des utilisateurs
func Register(w http.ResponseWriter, r *http.Request) {
	// Crée un nouveau compte et l'ajoute à la base de données
	// Vous devez ajouter la logique pour récupérer les données du formulaire et enregistrer l'utilisateur dans la base de données
	fmt.Fprintln(w, "Page Register")
}
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
*/