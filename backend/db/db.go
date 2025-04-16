package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func DBconnect() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Erreur ouverture connexion: %v", err))
	}

	// Vérifie la connexion
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Impossible de ping la DB: %v", err))
	}

	fmt.Println("Connexion à la base de données réussie")
}
