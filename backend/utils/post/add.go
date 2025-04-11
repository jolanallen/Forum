package post

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"Forum/backend/utils"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ToggleLikePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint64)
	postID, err := extractPostIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	post, err := getPostByID(postID)
	if err != nil {
		http.Error(w, "Post introuvable", http.StatusNotFound)
		return
	}

	hasLiked, err := hasUserLikedPost(userID, postID)
	if err != nil {
		http.Error(w, "Erreur lors de la vérification du like", http.StatusInternalServerError)
		return
	}

	if hasLiked {
		if err := removeLike(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors du retrait du like", http.StatusInternalServerError)
			return
		}
	} else {
		if err := addLike(userID, postID, &post); err != nil {
			http.Error(w, "Erreur lors de l'ajout du like", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

func extractPostIDFromURL(path string) (uint64, error) {
	postIDStr := strings.TrimPrefix(path, "/user/post/")
	postIDStr = strings.TrimSuffix(postIDStr, "/like")
	return strconv.ParseUint(postIDStr, 10, 64)
}

func getPostByID(postID uint64) (structs.Post, error) {
	var post structs.Post
	err := db.DB.First(&post, postID).Error
	return post, err
}

func hasUserLikedPost(userID, postID uint64) (bool, error) {
	var like structs.PostLike
	err := db.DB.Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}

func addLike(userID, postID uint64, post *structs.Post) error {
	newLike := structs.PostLike{
		UserID: userID,
		PostID: postID,
	}
	if err := db.DB.Create(&newLike).Error; err != nil {
		return err
	}
	post.PostLike++
	return db.DB.Save(post).Error
}

func removeLike(userID, postID uint64, post *structs.Post) error {
	if err := db.DB.Where("userID = ? AND postID = ?", userID, postID).Delete(&structs.PostLike{}).Error; err != nil {
		return err
	}
	if post.PostLike > 0 {
		post.PostLike--
	}
	return db.DB.Save(post).Error
}

func HandleCommentActions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		UserAddComment(w, r)
	case http.MethodPut:
		UserEditComment(w, r)
	case http.MethodDelete:
		UserDeleteComment(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}
func UserAddComment(w http.ResponseWriter, r *http.Request) {
	postID, err := extractPostIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	content := r.FormValue("comment")

	comment := structs.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
		Status:  "published",
		Visible: true,
	}

	if err := db.DB.Create(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}
func UserEditComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := extractCommentIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	var comment structs.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	updatedContent := r.FormValue("comment")
	comment.Content = updatedContent

	if err := db.DB.Save(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}
func UserDeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID, err := extractCommentIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "ID de commentaire invalide", http.StatusBadRequest)
		return
	}

	var comment structs.Comment
	if err := db.DB.First(&comment, commentID).Error; err != nil {
		http.Error(w, "Commentaire introuvable", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(uint64)
	if comment.UserID != userID {
		http.Error(w, "Non autorisé", http.StatusForbidden)
		return
	}

	if err := db.DB.Delete(&comment).Error; err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}

func extractCommentIDFromURL(path string) (uint64, error) {
	commentIDStr := strings.TrimPrefix(path, "/user/comment/")
	return strconv.ParseUint(commentIDStr, 10, 64)
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

func UserCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var categories []structs.Category
		if err := db.DB.Find(&categories).Error; err != nil {
			http.Error(w, "Erreur lors de la récupération des catégories", http.StatusInternalServerError)
			return
		}

		utils.Templates.ExecuteTemplate(w, "create_post.html", struct {
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
		PostKey:         postKey,
		PostComment:     content,
		UserID:          userID,
		ImageID:         imageID,
		CategoryID:      categoryID,
	}

	if err := db.DB.Create(&post).Error; err != nil {
		http.Error(w, "Erreur lors de la création du post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/forum/", http.StatusSeeOther)
}
