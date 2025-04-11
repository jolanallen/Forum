package db

import (
	"Forum/backend/structs"
	"log"
	"net/http"
	"text/template"
	//"github.com/mattn/go-sqlite3"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur recherché depuis les paramètres de la requête
	username := r.URL.Query().Get("username")

	// Chercher l'utilisateur dans la base de données
	var users []structs.User
	err := DB.Where("username LIKE ?", "%"+username+"%").Find(&users).Error
	if err != nil {
		log.Println("Erreur lors de la recherche des utilisateurs:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
		return
	}

	// Créer un template HTML et afficher les résultats
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		log.Println("Erreur lors du parsing du template:", err)
		http.Error(w, "Erreur de template", http.StatusInternalServerError)
		return
	}

	// Passer les utilisateurs trouvés au template
	err = tmpl.Execute(w, users)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template:", err)
		http.Error(w, "Erreur de serveur", http.StatusInternalServerError)
	}
}
