package handler

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
 // gestion de la requete verifie que l'utilaseur existe et que le ot de passe est le bon avec les focntion dans /utils/password.go
}

func Register(w http.ResponseWriter, r *http.Request) {
	//// crée un nouveau compte et l'ajoute a la bddd avec les fonction dans /db/ queries.go 
}
