package handler

import (
	"fmt"
	"net/http"
)

func UserEditProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//afiche le profil de l'utilisateur  en fait des requet a la bdd pour récuperer les info de l'utilisateur
		fmt.Fprintln(w, "Page de modification de l'utilisateur connecter")

	} else if r.Method == http.MethodPost {
		// Logique de mise à jour du profil
		//si et suelment si le profil qui demande a être a jour est celui de l'utilisateur qui fait la demande peut être introduire un middleware pour checker si le porfile est celui de l'utilisateur en cours ou pas  si oui ajouter les modfication dan sla bdd

	}
}

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Page creations post")

		// renvoi a la page de création post
	} else if r.Method == http.MethodPost {
		// Logique pour créer un post

	}
}

func UserLikePost(w http.ResponseWriter, r *http.Request) {
	// utilisation de l'id du post dans l'url et ajoute le like si et suelment si l'utilisateur actuelle na pas déja liker peut être introduire un middleware pour checker si liker ou pas  au post dans la base de donnée en utilisant queries.go
	fmt.Fprintln(w, "Page creations post")

}

func UserAddComment(w http.ResponseWriter, r *http.Request) {
	// utilisation de l'id du post dans l'url et ajoute le commentaire au post dans la base de donnée en utilisant queries.go
	fmt.Fprintln(w, "Page creations post")

}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "log out ")

	// Logique de déconnexion

}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page profils des autres user ")

}
