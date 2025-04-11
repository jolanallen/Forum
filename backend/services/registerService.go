package services

import (
	"Forum/backend/structs"
	"log"
	"net/http"
)

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

		if username == "" || email == "" || password == "" || confirmPassword == "" {
			http.Error(w, "Tous les champs doivent être remplis", http.StatusBadRequest)
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

		newUser := structs.User{
			UserUsername:     username,
			UserEmail:        email,
			UserPasswordHash: hashedPassword,
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
