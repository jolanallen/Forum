package db

import (
	"Forum/backend/structs"
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

	// Automigration pour créer ou mettre à jour les tables en fonction des structures
	err = DB.AutoMigrate(&structs.Topic{}, &structs.TopicDislike{}, &structs.TopicLike{})
	if err != nil {
		fmt.Println("Error during migration:", err)
		return
	}

	fmt.Println("Database connected and tables migrated successfully")
}
