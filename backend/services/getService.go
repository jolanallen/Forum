package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"fmt"
	"net/http"
	"time"
)

func GetPostByID(postID uint64) (structs.Post, error) {
	var post structs.Post
	err := db.DB.First(&post, postID).Error
	return post, err
}

func GetUserIDFromSession(r *http.Request) uint64 {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return 0
	}

	var userID uint64
	sessionType := cookie.Value[:5]

	switch sessionType {
	case "user_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "admin_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	case "guest_":
		fmt.Sscanf(cookie.Value[5:], "%d", &userID)
	default:
		return 0
	}

	return userID
}


func GetUserByID(userID uint64) (*structs.User, error) {
	var user structs.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// donne la structure de tout les poste en fonction de la categorie
func GetPostsByCategory(category string) ([]structs.Post, error) {
	var posts []structs.Post

	subQuery := db.DB. //// sous-requêtes qui permet de récupérer d'abord L'Id de la catégorie
				Table("categories").
				Select("categoryID").
				Where("LOWER(categoryName) = LOWER(?)", category)

	err := db.DB. /// récupére tout les post de la catégoroie correspondant a l'ID récupérer
			Where("categoryID = (?)", subQuery).
			Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetCommentByID(commentID uint64) (structs.Comment, error) {
	var comment structs.Comment
	err := db.DB.First(&comment, commentID).Error
	return comment, err
}

func GetUserByEmail(email string) (*structs.User, error) {
	fmt.Println(email)
	var user structs.User
	fmt.Println(user)
	result := db.DB.Model(&structs.User{}).Where("`userEmail` = ?", email).First(&user)

	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("blablaba")
	return &user, nil
}

func GetAdminDashboardData(adminID uint64) (*structs.AdminDashboardData, error) {
	var data structs.AdminDashboardData
	data.AdminID = adminID
	data.GeneratedAt = time.Now()

	var count int64

	err := db.DB.Model(&structs.User{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	data.TotalUsers = uint64(count)

	err = db.DB.Model(&structs.Post{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	data.TotalPosts = uint64(count)

	err = db.DB.Model(&structs.Comment{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	data.TotalComments = uint64(count)

	err = db.DB.Model(&structs.Guest{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	data.TotalGuests = uint64(count)

	return &data, nil
}
