package handler

import (
	"Forum/backend/db"
	"Forum/backend/services"
	"Forum/backend/structs"
	"database/sql"
	"log"
	"net/http"
)

// handleError is a helper function to log and return HTTP errors
func handleError(w http.ResponseWriter, err error, message string) {
	log.Println(message, err)
	http.Error(w, message, http.StatusInternalServerError)
}

// GuestHome handles the rendering of the home page for both authenticated and unauthenticated users
func GuestHome(w http.ResponseWriter, r *http.Request) {
	userID := services.GetUserIDFromSession(r)
	isAuthenticated := userID != 0

	// Fetch all posts with their authors using raw SQL
	var posts []structs.Post
	rows, err := db.DB.Query(`
		SELECT posts.postID, posts.categoryID, posts.postKey, posts.imageID, posts.postComment, posts.postLike, posts.createdAt, 
		       users.username
		FROM posts
		JOIN users ON posts.userID = users.userID
	`)
	if err != nil {
		handleError(w, err, "Error while fetching posts")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var username string
		if err := rows.Scan(&post.PostID, &post.CategoryID, &post.PostKey, &post.ImageID, &post.PostComment, &post.PostLike, &post.PostCreatedAt, &username); err != nil {
			handleError(w, err, "Error while scanning post data")
			return
		}
		post.UserUsername = username
		posts = append(posts, post)
	}

	// Fetch all categories
	var categories []structs.Category
	rows, err = db.DB.Query("SELECT categoryID, categoryName FROM categories")
	if err != nil {
		handleError(w, err, "Error while fetching categories")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			handleError(w, err, "Error while scanning category data")
			return
		}
		categories = append(categories, category)
	}

	// If the user is authenticated, retrieve their user data
	var user structs.User
	if isAuthenticated {
		row := db.DB.QueryRow("SELECT userID, userUsername FROM users WHERE userID = $1", userID)
		if err := row.Scan(&user.UserID, &user.UserUsername); err != nil {
			if err == sql.ErrNoRows {
				handleError(w, err, "User not found")
			} else {
				handleError(w, err, "Error while fetching user")
			}
			return
		}
	}

	// Render the home page template with posts, categories, and user info
	services.RenderTemplate(w, "forum/home.html", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
		User            structs.User
		ActivePage      string
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
		User:            user,
		ActivePage:      "home",
	})
}

// CategoryHack handles displaying posts from the "hack" category
func CategoryHack(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("hack")
	if err != nil {
		handleError(w, err, "Error while fetching Hack posts")
		return
	}
	services.RenderTemplate(w, "/categories_post.html", posts)
}

// CategoryProg handles displaying posts from the "prog" category
func CategoryProg(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("prog")
	if err != nil {
		handleError(w, err, "Error while fetching Prog posts")
		return
	}
	services.RenderTemplate(w, "/categories_post.html", posts)
}

// CategoryNews handles displaying posts from the "news" category
func CategoryNews(w http.ResponseWriter, r *http.Request) {
	posts, err := services.GetPostsByCategory("news")
	if err != nil {
		handleError(w, err, "Error while fetching News posts")
		return
	}
	services.RenderTemplate(w, "/categories_post.html", posts)
}

// SearchPseudo handles searching for posts by keyword or username
func SearchPseudo(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("query")
	posts, err := services.SearchPosts(searchQuery)
	if err != nil {
		handleError(w, err, "Error during post search")
		return
	}
	services.RenderTemplate(w, "/search.html", posts)
}

// AboutForum renders the static "About" page
func AboutForum(w http.ResponseWriter, r *http.Request) {
	services.RenderTemplate(w, "/about.html", nil)
}
