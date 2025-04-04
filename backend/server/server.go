package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"Forum/backend/db"
	"Forum/backend/utils"
)

func Server() {
	// Lire les variables d'environnement pour les certificats
	certFile := os.Getenv("CERT_PATH")
	keyFile := os.Getenv("KEY_PATH")

	// Vérifier si les variables d'environnement sont définies
	if certFile == "" || keyFile == "" {
		log.Fatal("❌ Les variables d'environnement CERT_PATH ou KEY_PATH ne sont pas définies.")
	}

	InitRoutes()
	db.DBconnect()
	fmt.Println("https://localhost:443")

	err := http.ListenAndServeTLS(":443", certFile, keyFile, utils.F.MainRouter)
	if err != nil {
		log.Fatal(err)
	}
}
