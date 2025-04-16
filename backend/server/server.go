package server

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Server() {
	// Read the environment variables for the SSL certificate paths
	certFile := os.Getenv("CERT_PATH") // Get the certificate file path from the environment variable
	keyFile := os.Getenv("KEY_PATH")   // Get the key file path from the environment variable

	// Check if the environment variables are set
	if certFile == "" || keyFile == "" {
		log.Fatal("❌ The environment variables CERT_PATH or KEY_PATH are not set.") // Exit if paths are not set
	}

	// Initialize routes for the server
	InitRoutes()

	// Connect to the database
	db.DBconnect()

	// Print the URL of the server to the console
	fmt.Println("https://localhost:443/forum/")

	// Start the server with TLS (HTTPS) using the provided certificate and key files
	err := http.ListenAndServeTLS(":443", certFile, keyFile, services.F.MainRouter)
	if err != nil {
		log.Fatal(err) // Log and terminate the program if there is an error starting the server
	}
}
