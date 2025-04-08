package handler

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	var admins []structs.Admin
	result := db.DB.Find(&admins) // Récupérer tous les admins
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(admins)
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
