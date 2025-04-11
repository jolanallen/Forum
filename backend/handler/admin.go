package handler

import (
	"fmt"
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {

//////////////dans adminService.go/////////////////////

}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un utilisateur

//////////////dans adminService.go/////////////////////

	fmt.Fprintln(w, "Page admin delete user")
}

func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	// logic de supression d'un commentaire

//////////////dans adminService.go/////////////////////

	fmt.Fprintln(w, "Page admin delete comment")
}

func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un post

//////////////dans adminService.go/////////////////////

	fmt.Fprintln(w, "Page admin delete post")
}
