package handler

import (
	"fmt"
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {

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
