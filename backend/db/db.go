package db

import (
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBconnect() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/forum?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}
	fmt.Println("Database connected successfully")

	// Charger et exécuter le fichier de migration
	err = runMigrationScript("databases/forum.sql")
	if err != nil {
		log.Println("Error executing migration script:", err)
		return
	}

	fmt.Println("Migration script executed successfully")
}

func runMigrationScript(filePath string) error {
	// Lire le contenu du fichier SQL
	sqlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %v", err)
	}

	// Exécuter le script SQL
	sql := string(sqlBytes)
	if err := DB.Exec(sql).Error; err != nil {
		return fmt.Errorf("failed to execute SQL script: %v", err)
	}

	return nil
}

