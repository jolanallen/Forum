package server

import (
	"fmt"
	"log"
	"net/http"
	
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenue sur mon serveur HTTPS !"))
}

func Server() {
	
	certFile := "backend/server/cert.pem"
	keyFile := "backend/server/key.pem"

	http.HandleFunc("/", handler)

	fmt.Println("https://localhost:443")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}
