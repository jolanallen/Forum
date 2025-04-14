package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var categories []structs.Category
		if err := db.DB.Find(&categories).Error; err != nil {
			http.Error(w, "Erreur lors de la récupération des catégories", http.StatusInternalServerError)
			return
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

	imageID, err := services.HandleImageUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := structs.Post{
		PostKey:     postKey,
		PostComment: content,
		UserID:      userID,
		ImageID:     imageID,
		CategoryID:  categoryID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
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
		if err := services.RemoveLikeFromComment(userID, commentID, &comment); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := services.AddLikeToComment(userID, commentID, &comment); err != nil {
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
		if err := services.RemoveLikeFromPost(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := services.AddLikeToPost(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/forum", http.StatusSeeOther)
}
func UserEditProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	vars := mux.Vars(r)
	// Si nécessaire, tu peux récupérer l'id aussi de cette manière
	_ = vars["id"] // L'id dans l'URL est utilisé ici pour la vérification si nécessaire

	if r.Method == http.MethodGet {
		user, err := services.GetUserByID(userID)
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
			imageID, err := services.ValidateImage(file, nil)
			if err != nil {
				http.Error(w, "Erreur de validation de l'image : "+err.Error(), http.StatusBadRequest)
				return
			}
			user.UserProfilePicture = &imageID.ImageID
		}

		if err := UpdateUser(user); err != nil {
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

func UpdateUser(user *structs.User) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *structs.User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
