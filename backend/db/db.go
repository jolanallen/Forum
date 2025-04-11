package db

import (
	"fmt"

	"gorm.io/driver/mysql" // Utilisation du driver MySQL de GORM v2
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBconnect initialise la connexion à la base de données MySQL
func DBconnect() {
	var err error

	// DSN pour la connexion à MySQL (remplacer avec les bonnes informations)
	dsn := "root:root@tcp(localhost:3306)/forum?parseTime=true"

	// Connexion à MySQL avec GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Database connected and tables migrated successfully")
}
