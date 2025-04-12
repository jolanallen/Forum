package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	
	"net/http"
	"strconv"
)
//pdp
func UserEditProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)

	if r.Method == http.MethodGet {
		user, err := GetUserByID(userID)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des informations du profil", http.StatusInternalServerError)
			return
		}
		Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", user)
	} else if r.Method == http.MethodPost {
		newUsername := r.FormValue("userUsername")
		newEmail := r.FormValue("userEmail")
		newPassword := r.FormValue("userPassword")
		user, err := GetUserByID(userID)
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
			hashedPassword, err := HashPassword(newPassword)
			if err != nil {
				http.Error(w, "Erreur lors du hachage du mot de passe", http.StatusInternalServerError)
				return
			}
			user.UserPasswordHash = hashedPassword
		}

		if file, _, err := r.FormFile("userProfilePicture"); err == nil {
			imageID, err := validateImage(file, nil)
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

		http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
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

	user, err := GetUserByID(userID)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}
	Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", user)
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