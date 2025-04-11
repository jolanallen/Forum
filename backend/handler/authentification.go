package handler

import (
	"fmt"
	"net/http"
)

// Assurez-vous d'importer le package contenant les structures (par exemple, models)
// Login gère la logique de connexion de l'utilisateur
func Login(w http.ResponseWriter, r *http.Request) {

}

// Register gère l'inscription des utilisateurs
func Register(w http.ResponseWriter, r *http.Request) {
	// Crée un nouveau compte et l'ajoute à la base de données
	// Vous devez ajouter la logique pour récupérer les données du formulaire et enregistrer l'utilisateur dans la base de données
	fmt.Fprintln(w, "Page Register")
}
