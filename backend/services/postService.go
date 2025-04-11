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

	"github.com/google/uuid"
)

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var categories []structs.Category
		if err := db.DB.Find(&categories).Error; err != nil {
			http.Error(w, "Erreur lors de la récupération des catégories", http.StatusInternalServerError)
			return
		}

		Templates.ExecuteTemplate(w, "create_post.html", struct {
			Categories []structs.Category
		}{
			Categories: categories,
		})
		return
	}

	userID := r.Context().Value("userID").(uint64)
	postKey := uuid.New().String()

	content, categoryID, err := parseFormValues(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imageID, err := handleImageUpload(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post := structs.Post{
		PostKey:     postKey,
		PostComment: content,
		UserID:      userID,
		ImageID:     imageID,
		CategoryID:  categoryID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

// Valide l’image (format et taille)
func validateImage(file multipart.File, header *multipart.FileHeader) (*structs.Image, error) {
	const maxSize = 20 << 20 // 20 MB

	if header.Size > maxSize {
		return nil, fmt.Errorf("Image trop lourde (max 20MB)")
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
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

// Parse les champs du formulaire (content, description, category_id)
func parseFormValues(r *http.Request) (string, uint64, error) {
	content := r.FormValue("content")
	categoryIDStr := r.FormValue("category_id")

	if categoryIDStr == "" {
		return "", 0, fmt.Errorf("Catégorie obligatoire")
	}

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("ID de catégorie invalide")
	}

	return content, categoryID, nil
}

// Gère l'upload de l’image (renvoie l’image ID ou nil)
func handleImageUpload(r *http.Request) (*uint64, error) {
	file, header, err := r.FormFile("image")
	if err != nil {
		return nil, nil // pas d'image envoyée = pas une erreur
	}
	defer file.Close()

	image, err := validateImage(file, header)
	if err != nil {
		return nil, err
	}

	if err := db.DB.Create(image).Error; err != nil {
		return nil, fmt.Errorf("Erreur DB image")
	}

	return &image.ImageID, nil
}
