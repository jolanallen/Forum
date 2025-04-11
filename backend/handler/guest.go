package handler

import (
	"fmt"
	"net/http"
)

func GuestHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page d'accueil des invités become")
	
//////////////dans guestService.go/////////////////////

}

func GuestHack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Hack pour les invités")
	
//////////////dans guestService.go/////////////////////

}

func GuestProg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page Programme pour les invités")
	
//////////////dans guestService.go/////////////////////

}

func GuestNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page des actualités pour les invités")

//////////////dans guestService.go/////////////////////

}

//je vois vraiment pas ce que je peux faire avec ça fiston
//a la limite search bar pour un profile, nan j'ai déjà searchHandler dans searchService.go
// bah écoute, j'ai mis recherche dans les commentaire des posts
func GuestSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page de recherche pour les invités")
}

//et ça encore moins, genre about quoi??????????????
func GuestAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Page À propos pour les invités")
}
