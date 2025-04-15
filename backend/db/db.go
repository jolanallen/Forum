package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"Forum/backend/structs" // Assurez-vous que ce chemin est correct
)

var DB *gorm.DB

// DBconnect initialise la connexion à la base de données et effectue la migration.
func DBconnect() {
	var err error

	// Chaîne de connexion à la base de données
	dsn := "root:root@tcp(localhost:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Active les logs GORM
	})
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données:", err)
		return
	}
	DB.Migrator().DropTable(
		&structs.AdminDashboardData{},
		&structs.CommentLike{},
		&structs.PostLike{},
		&structs.Comment{},
		&structs.Post{},
		&structs.Category{},
		&structs.Guest{},
		&structs.Admin{},
		&structs.SessionUser{},
		&structs.SessionAdmin{},
		&structs.SessionGuest{},
		&structs.User{},
		&structs.Image{},
	)

	// Désactiver temporairement les vérifications de clés étrangères
	DB.Exec("SET foreign_key_checks = 0;")

	// Migration des tables dans le bon ordre
	if err := DB.AutoMigrate(
		&structs.Image{},
		&structs.User{},
		&structs.Category{},
		&structs.SessionUser{},
		&structs.AdminDashboardData{},
		&structs.Admin{},
		&structs.Guest{},
		&structs.SessionAdmin{},
		&structs.SessionGuest{},
		&structs.Comment{}, // Table `comments` doit être créée avant `posts`
		&structs.Post{},    // Table `posts` doit être créée après `comments`
		&structs.CommentLike{},
		&structs.PostLike{},
	); err != nil {
		log.Fatal("Erreur lors de la migration des tables: ", err)
	}

	// Réactiver les vérifications de clés étrangères
	DB.Exec("SET foreign_key_checks = 1;")

	// Message de succès
	fmt.Println("Base de données connectée et migration effectuée avec succès")
}
