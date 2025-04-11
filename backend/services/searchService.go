package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"net/http"
	"text/template"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur recherché depuis les paramètres de la requête
	username := r.URL.Query().Get("username")

	var users []structs.User
	if err := db.DB.Where("username LIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
		//erreur 500
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	//FAUT FAIRE EN SORTE QU4IL RENVOIE VERS LE TEMPLATE DU PROFIL DE L4USER CHERCHER

	//OU ALORS ON LES PARSE DIRECT

	// Créer un template HTML et afficher les résultats
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}

	// Passer les utilisateurs trouvés au template
	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}
