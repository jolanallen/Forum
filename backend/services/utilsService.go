package services

import (
	"Forum/backend/structs"
	"log"
	"text/template"
)

var Templates *template.Template

var F = &structs.Forum{}


// InitTemplates initialise les templates HTML
func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du parsing des templates:", err)
	}
}
