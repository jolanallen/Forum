package handler

import (
	"fmt"
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	// affiche la page dan le {{contenent de la page html le dasborad avec tout les élément }} il y aura trois dasbord a fair en focntion de la section si on est dans hack news ou prog 
	fmt.Fprintln(w, "Page admin dashboard")
}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un utilisateur 
	fmt.Fprintln(w, "Page admin delete user")
}

func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	// logic de supression d'un commentaire 
	fmt.Fprintln(w, "Page admin delete comment")
}

func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un post
	fmt.Fprintln(w, "Page admin delete post")
}
