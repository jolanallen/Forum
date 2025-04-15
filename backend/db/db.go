package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Forum/backend/structs" // Assurez-vous d'importer les structs appropriées
)

var DB *gorm.DB

// DBconnect initialise la connexion à la base de données et effectue la migration.
func DBconnect() {
	var err error

	// Chaîne de connexion à la base de données (ajuster selon ton environnement)
	dsn := "root:root@tcp(172.30.0.2:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données:", err)
		return
	}

	// Effectuer la migration des tables en utilisant AutoMigrate
	if err := DB.AutoMigrate(&structs.Post{}, &structs.Comment{}, &structs.Like{}); err != nil {
		fmt.Println("Erreur lors de la migration des tables:", err)
		return
	}
	
	// Message de succès
	fmt.Println("Base de données connectée et migration effectuée avec succès")
}
