package handler

import (
	
	"net/http"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	// affiche la page dan le {{contenent de la page html le dasborad avec tout les élément }} il y aura trois dasbord a fair en focntion de la section si on est dans hack news ou prog 
}

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un utilisateur 
}

func AdminDeleteComment(w http.ResponseWriter, r *http.Request) {
	// logic de supression d'un commentaire 
}

func AdminDeletePost(w http.ResponseWriter, r *http.Request) {
	//logic de suppression d'un post
}
