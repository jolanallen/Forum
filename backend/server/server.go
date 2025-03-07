package server

import (
	"fmt"
	"log"
	"net/http"
	
)
// ici seule les route pour la page principales 
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenue sur mon serveur HTTPS !"))
}

func Server() {
	
	certFile := "backend/server/ssl_tls/cert.pem"
	keyFile := "backend/server/ssl_tls/key.pem"

	http.HandleFunc("/", handler)

	fmt.Println("https://localhost:443")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}
