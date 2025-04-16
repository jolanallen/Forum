package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"strconv"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"database/sql"
)

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Récupérer les catégories via une requête SQL brute
		var categories []structs.Category
		rows, err := db.DB.Query("SELECT categoryID, categoryName FROM categories")
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des catégories", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var category structs.Category
			if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
				http.Error(w, "Erreur lors du traitement des catégories", http.StatusInternalServerError)
				return
			}
			categories = append(categories, category)
		}

		services.RenderTemplate(w, "BoyWithUke_Prairies", struct {
			Categories []structs.Category
		}{
			Categories: categories,
		})
		return
	}

	userID := r.Context().Value("userID").(uint64)
	postKey := uuid.New().String()

	content, categoryID, err := services.ParseFormValues(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Handle image upload
	imageID, err := services.HandleImageUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the post
	query := "INSERT INTO posts (postKey, postComment, userID, imageID, categoryID) VALUES (?, ?, ?, ?, ?)"
	_, err = db.DB.Exec(query, postKey, content, userID, imageID, categoryID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la création du post : %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/post-created", http.StatusSeeOther) // Rediriger vers une URL de confirmation
}

func ToggleLikeComment(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	commentIDStr := vars["id"]

	commentID, err := strconv.ParseUint(commentIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	comment, err := services.GetCommentByID(commentID)
	if err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	hasLiked, err := services.HasUserLikedComment(userID, commentID)
	if err != nil {
		http.Error(w, "Erreur lors de la vérification du like", http.StatusInternalServerError)
		return
	}

	if hasLiked {
		// Retirer le like
		query := "DELETE FROM comment_likes WHERE userID = ? AND commentID = ?"
		_, err := db.DB.Exec(query, userID, commentID)
		if err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		// Ajouter un like
		query := "INSERT INTO comment_likes (userID, commentID) VALUES (?, ?)"
		_, err := db.DB.Exec(query, userID, commentID)
		if err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

func ToggleLikePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	postIDStr := vars["id"]

	postID, err := strconv.ParseUint(postIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de post invalide", http.StatusBadRequest)
		return
	}

	post, err := services.GetPostByID(postID)
	if err != nil {
		http.Error(w, "Post introuvable", http.StatusNotFound)
		return
	}

	hasLiked, err := services.HasUserLikedPost(userID, postID)
	if err != nil {
		http.Error(w, "Erreur lors de la vérification du like", http.StatusInternalServerError)
		return
	}

	if hasLiked {
		// Retirer le like
		query := "DELETE FROM post_likes WHERE userID = ? AND postID = ?"
		_, err := db.DB.Exec(query, userID, postID)
		if err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		// Ajouter un like
		query := "INSERT INTO post_likes (userID, postID) VALUES (?, ?)"
		_, err := db.DB.Exec(query, userID, postID)
		if err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}

func UserEditProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	_ = vars["id"]

	if r.Method == http.MethodGet {
		query := "SELECT userID, userUsername, userEmail, userProfilePicture FROM users WHERE userID = ?"
		var user structs.User
		err := db.DB.QueryRow(query, userID).Scan(&user.UserID, &user.UserUsername, &user.UserEmail, &user.UserProfilePicture)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations du profil", http.StatusInternalServerError)
			return
		}
		services.RenderTemplate(w, "profile_edit.html", user)
	} else if r.Method == http.MethodPost {
		newUsername := r.FormValue("userUsername")
		newEmail := r.FormValue("userEmail")
		newPassword := r.FormValue("userPassword")
		user, err := services.GetUserByID(userID)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération de l'utilisateur", http.StatusInternalServerError)
			return
		}
		if newUsername != "" {
			user.UserUsername = newUsername
		}
		if newEmail != "" {
			user.UserEmail = newEmail
		}
		if newPassword != "" {
			hashedPassword, err := services.HashPassword(newPassword)
			if err != nil {
				http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
				return
			}
			user.UserPasswordHash = hashedPassword
		}

		if file, _, err := r.FormFile("userProfilePicture"); err == nil {
			image, err := services.ValidateImage(file, nil)
			if err != nil {
				http.Error(w, "Erreur de validation de l'image : "+err.Error(), http.StatusBadRequest)
				return
			}
			user.UserProfilePicture = image.ImageID
		}

		query := "UPDATE users SET userUsername = ?, userEmail = ?, userPasswordHash = ?, userProfilePicture = ? WHERE userID = ?"
		_, err = db.DB.Exec(query, user.UserUsername, user.UserEmail, user.UserPasswordHash, user.UserProfilePicture, user.UserID)
		if err != nil {
			http.Error(w, "Erreur lors de la mise à jour du profil", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Path[len("/user/"):]

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "ID d'utilisateur invalide", http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}
	services.RenderTemplate(w, "BoyWithUke_Prairies", user)
}
