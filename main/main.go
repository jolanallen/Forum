package main

import (
	// Importer le package server

	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Détails de la connexion
	dsn := "root:root@tcp(127.0.0.1:3306)/forum"

	// Ouvrir la connexion à la base de données
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tester la connexion
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connexion à la base de données réussie!")

	// Exemple de récupération de données
	rows, err := db.Query("SELECT id FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Println("Connsbhhhsie!")
	for rows.Next() {
		var id int
		var nom string
		fmt.Println("Connssgfjhgihggie!")
		if err := rows.Scan(&id, &nom); err != nil {
			log.Fatal(err)
			fmt.Println("Connssie!")
		}
		fmt.Printf("ID: %d, Nom: %s\n", id, nom)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	server.Server() // Démarrer le serveur HTTPS
}
*/
