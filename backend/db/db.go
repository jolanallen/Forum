package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql" // Import du pilote MySQL
)

var DB *sql.DB

func DBconnect() {
	var err error
	// Remplacez cette ligne par votre chaîne de connexion MySQL
	// Exemple : "root:password@tcp(127.0.0.1:3306)/forum"
	DB, err = sql.Open("mysql", "root:password@tcp(forum-mysql:3306)/forum")
	if err != nil {
		log.Fatal(err)
	}
	
}