package handler

import (
	"fmt"
	"net/http"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités")
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
