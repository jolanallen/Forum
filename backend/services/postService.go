package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func ValidateImage(file multipart.File, header *multipart.FileHeader) (*structs.Image, error) {
	const maxSize = 20 << 20 // 20MB

	if header.Size > maxSize {
		return nil, fmt.Errorf("Image trop lourde (max 20MB)")
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".png" && ext != ".gif" {
		return nil, fmt.Errorf("Format d'image non supporté")
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, fmt.Errorf("Erreur de lecture du fichier")
	}

	image := &structs.Image{
		Filename: header.Filename,
		Data:     buf.Bytes(),
		URL:      "/images/" + header.Filename,
	}

	return image, nil
}

func ParseFormValues(r *http.Request) (string, uint64, error) {
	content := r.FormValue("content")
	categoryIDStr := r.FormValue("categoriesID")

	if categoryIDStr == "" {
		return "", 0, fmt.Errorf("Catégorie obligatoire")
	}

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("ID de catégorie invalide")
	}

	return content, categoryID, nil
}

func HandleImageUpload(r *http.Request) (uint64, error) {
	file, header, err := r.FormFile("image")
	if err != nil {
		return 0, fmt.Errorf("Erreur lors de la récupération du fichier: %v", err)
	}
	defer file.Close()

	image, err := ValidateImage(file, header)
	if err != nil {
		return 0, fmt.Errorf("Image invalide: %v", err)
	}

	// Insertion SQL manuelle (sans GORM)
	query := `INSERT INTO images (filename, data, url) VALUES (?, ?, ?)`
	result, err := db.DB.Exec(query, image.Filename, image.Data, image.URL)
	if err != nil {
		return 0, fmt.Errorf("Erreur base de données: %v", err)
	}

	imageID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Erreur récupération ID image: %v", err)
	}

	return uint64(imageID), nil
}
