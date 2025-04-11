package utils

import (
	"Forum/backend/structs"
	"crypto/rand"
	"encoding/hex"
	"html/template"
	"log"
)

var Templates *template.Template

func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du parsing des templates:", err)
	}
}

var F = &structs.Forum{} ///variables global ( à voir ce que je voulais en faire)


// création d'un token de sessions aléatoire de byte en hexa
func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
