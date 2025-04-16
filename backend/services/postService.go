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

// ValidateImage checks if the uploaded image is valid, ensuring it meets size and format requirements.
func ValidateImage(file multipart.File, header *multipart.FileHeader) (*structs.Image, error) {
	const maxSize = 20 << 20 // 20MB

	// Check if the image exceeds the maximum size limit
	if header.Size > maxSize {
		return nil, fmt.Errorf("Image too large (max 20MB)")
	}

	// Validate the image format (only jpg, png, and gif are allowed)
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".png" && ext != ".gif" {
		return nil, fmt.Errorf("Unsupported image format")
	}

	// Copy the image content into a buffer for further processing
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, fmt.Errorf("Error reading the file")
	}

	// Create an Image struct containing the image's data and URL
	image := &structs.Image{
		Filename: header.Filename,
		Data:     buf.Bytes(),
		URL:      "/images/" + header.Filename,
	}

	return image, nil
}

// ParseFormValues extracts and validates the form values for content and categoryID.
func ParseFormValues(r *http.Request) (string, uint64, error) {
	// Get the content and categoryID from the form data
	content := r.FormValue("content")
	categoryIDStr := r.FormValue("categoriesID")

	// Check if categoryID is provided
	if categoryIDStr == "" {
		return "", 0, fmt.Errorf("Category is required")
	}

	// Parse the categoryID from string to uint64
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		return "", 0, fmt.Errorf("Invalid category ID")
	}

	return content, categoryID, nil
}

// HandleImageUpload processes the image upload, validates it, and stores it in the database.
func HandleImageUpload(r *http.Request) (uint64, error) {
	// Get the uploaded image file and its header
	file, header, err := r.FormFile("image")
	if err != nil {
		return 0, fmt.Errorf("Error retrieving the file: %v", err)
	}
	defer file.Close()

	// Validate the uploaded image
	image, err := ValidateImage(file, header)
	if err != nil {
		return 0, fmt.Errorf("Invalid image: %v", err)
	}

	// Manually insert the image data into the database
	query := `INSERT INTO images (filename, data, url) VALUES (?, ?, ?)`
	result, err := db.DB.Exec(query, image.Filename, image.Data, image.URL)
	if err != nil {
		return 0, fmt.Errorf("Database error: %v", err)
	}

	// Retrieve the last inserted image ID from the database
	imageID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error retrieving image ID: %v", err)
	}

	return uint64(imageID), nil
}
