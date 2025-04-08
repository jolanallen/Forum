package handler

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"log"
	"net/http"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités b")
}

func GuestGuest(w http.ResponseWriter, r *http.Request) {
	var guests []structs.Guest
	result := db.DB.Find(&guests) // Récupérer tous les guests
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(guests)
}

func GuestHack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Hack pour les invités")
}

func GuestProg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Programme pour les invités")
}

func GuestNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page des actualités pour les invités")
}

func GuestSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page de recherche pour les invités")
}

func GuestAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page À propos pour les invités")
}
