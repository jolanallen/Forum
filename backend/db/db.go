package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBconnect() {
	var err error

	// Mets ici l'adresse IP de ton conteneur MySQL
	dsn := "root:root@tcp(172.30.0.2:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Erreur ouverture connexion: %v", err))
	}

	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Impossible de ping la DB: %v", err))
	}

	fmt.Println("Connexion à la base de données réussie")
}
