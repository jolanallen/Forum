package utils

import (
	"Forum/backend/structs"
	"bytes"
	"io"
	"net/http"

	"gorm.io/gorm"
)

/*lien pour comprendre gorm : https://pkg.go.dev/gorm.io/gorm */

func downloadImage(url string) ([]byte, string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	parts := bytes.Split([]byte(url), []byte("/"))
	filename := string(parts[len(parts)-1])

	return data, filename, nil
}

func createPostWithImage(db *gorm.DB, postKey, imageURL, commentaire string, creatorID, userID uint) error {
	data, filename, err := downloadImage(imageURL)
	if err != nil {
		return err
	}

	post := structs.Post{
		PostKey: postKey,
		Image: structs.Image{
			URL:      imageURL,
			Filename: filename,
			Data:     data,
		},
		Commentaire: commentaire,
		CreatorID:   creatorID,
		UserID:      userID,
	}

	result := db.Create(&post)
	return result.Error
}
