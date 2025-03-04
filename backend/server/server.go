package server


import (
	"net/http"
	"log"

)
// chemin accés au certificat ssl tls et a ca clef
var  Certfile string = "/ssl_tls/cert.cert"
var keyfile string = "/ssl_tls/default.key"

func Server() {
	err := http.ListenAndServeTLS(":8080", keyfile, Certfile, nil) // serveur web sur le port 8080 avec un certificat ssl tls et sa clef 
	if err != nil {
			log.Fatal(err)
	}

}