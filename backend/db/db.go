package db

import (
    "database/sql"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBconnect() {
    var err error
    var retries int

    // Connexion à MySQL (remplacer avec les bonnes informations)
    dsn := "root:root@tcp(forum-mysql:3306)/forum?parseTime=true"

    // Attendre 20 secondes avant de commencer à tenter la connexion
    log.Println("⏳ Attente de 20 secondes avant d'essayer de se connecter à la BDD...")
    time.Sleep(20 * time.Second)

    // Tentatives de connexion avec un délai entre chaque essai
    for retries = 0; retries < 10; retries++ {
        DB, err = sql.Open("mysql", dsn)
        if err != nil {
            log.Printf("Erreur ouverture BDD, tentative %d : %v\n", retries+1, err)
            time.Sleep(5 * time.Second) // Attendre 5 secondes avant de réessayer
            continue
        }

        DB.SetConnMaxLifetime(time.Minute * 3)
        DB.SetMaxOpenConns(10)
        DB.SetMaxIdleConns(5)

        if err = DB.Ping(); err != nil {
            log.Printf("Erreur connexion BDD, tentative %d : %v\n", retries+1, err)
            time.Sleep(5 * time.Second) // Attendre encore avant de réessayer
            continue
        }

        log.Println("Connexionréussie !")
        break // Si tout va bien, sortir de la boucle
    }

}
