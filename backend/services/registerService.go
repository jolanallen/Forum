package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"log"
	"net/http"
)
//utilisé dans backend\handler\authentification.go
//enregistrer un nouvel utilisateur
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// j'imagine que c'est pour register.html mais dans le doute je mettrais prairies
		Templates.ExecuteTemplate(w, "BoyWithUke_Prairies", nil)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("userUsername")
		email := r.FormValue("userEmail")
		password := r.FormValue("userPassword")
		confirmPassword := r.FormValue("confirm_password")

		err := CheckRegistrationForm(username, email, password, confirmPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if password != confirmPassword {
			http.Error(w, "Les mots de passe ne correspondent pas", http.StatusBadRequest)
			return
		}

		existingUser, err := GetUserByEmail(email)
		if err != nil {
			http.Error(w, "Erreur lors de la recherche de l'email", http.StatusInternalServerError)
			return
		}
		if existingUser != nil {
			http.Error(w, "Un compte avec cet email existe déjà", http.StatusBadRequest)
			return
		}

		hashedPassword, err := HashPassword(password)
		if err != nil {
			log.Println("Erreur lors du hachage du mot de passe:", err)
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}
		var userProfileImageID *uint64
		userProfileImageID, err = handleImageUpload(r)
		if err != nil {
			log.Println("Erreur d'upload d'image, utilisation de l'image par défaut")
			defaultImage := structs.Image{
				Filename: "default.png",
				URL:      "/images/default.png",
			}

			if err := db.DB.Create(&defaultImage).Error; err != nil {
				log.Println("Erreur lors de l'ajout de l'image par défaut:", err)
				http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
				return
			}
			userProfileImageID = &defaultImage.ImageID
		}
		newUser := structs.User{
			UserUsername:       username,
			UserEmail:          email,
			UserPasswordHash:   hashedPassword,
			UserProfilePicture: userProfileImageID,
		}

		if err := CreateUser(&newUser); err != nil {
			log.Println("Erreur lors de la création de l'utilisateur:", err)
			http.Error(w, "Erreur lors de l'inscription", http.StatusInternalServerError)
			return
		}
		///je ne sais tjr pas la bonne route
		http.Redirect(w, r, "BoyWithUke_Prairies", http.StatusSeeOther)
	}
}
