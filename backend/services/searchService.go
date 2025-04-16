package services

import (
	"Forum/backend/db"
	"Forum/backend/structs"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// SearchHandler handles the search requests for users based on their username.
// It validates the form, queries the database, and either redirects or displays search results.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Retrieve the username from the form
	username := r.FormValue("userUsername")
	// Query the database to find users with similar usernames
	rows, err := db.DB.Query("SELECT userID, userUsername FROM users WHERE userUsername LIKE ?", "%"+username+"%")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("Query error:", err)
		return
	}
	defer rows.Close()

	// Collect the found users
	var users []structs.User
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.UserID, &user.UserUsername); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		users = append(users, user)
	}

	// If exactly one user is found, redirect to their profile
	if len(users) == 1 {
		http.Redirect(w, r, "/profile/"+users[0].UserUsername, http.StatusSeeOther)
		return
	}

	// If multiple or no users are found, display the results
	tmpl, err := template.ParseFiles("templates/search_results.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

// SearchPosts performs a search on posts by their content (post comment).
// It returns a slice of posts that match the search query.
func SearchPosts(query string) ([]structs.Post, error) {
	// Query the database for posts that match the search query
	rows, err := db.DB.Query("SELECT postID, postComment, userID, postKey, imageID, postLike, categoryID, createdAt FROM posts WHERE postComment LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Collect the matching posts
	var posts []structs.Post
	for rows.Next() {
		var p structs.Post
		err := rows.Scan(&p.PostID, &p.PostComment, &p.UserID, &p.PostKey, &p.ImageID, &p.PostLike, &p.CategoryID, &p.PostCreatedAt)
		if err != nil {
			log.Println("Post scan error:", err)
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// FilterPostsByCategory handles the filtering of posts by their category.
// It validates the category ID and retrieves the relevant posts from the database.
func FilterPostsByCategory(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Retrieve and validate the category ID from the form
	categoryIDStr := r.FormValue("categoriesID")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	// Query the database to get posts for the selected category
	postRows, err := db.DB.Query(`SELECT postID, postComment, userID, postKey, imageID, postLike, categoryID, createdAt FROM posts WHERE categoryID = ?`, categoryID)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		log.Println("Error retrieving posts:", err)
		return
	}
	defer postRows.Close()

	// Collect the posts for the specified category
	var posts []structs.Post
	for postRows.Next() {
		var post structs.Post
		if err := postRows.Scan(&post.PostID, &post.PostComment, &post.UserID, &post.PostKey, &post.ImageID, &post.PostLike, &post.CategoryID, &post.PostCreatedAt); err != nil {
			log.Println("Post scan error:", err)
			continue
		}
		posts = append(posts, post)
	}

	// Check if the user is authenticated
	userID := GetUserIDFromSession(r)
	isAuthenticated := userID != 0

	// Retrieve the categories from the database
	catRows, err := db.DB.Query("SELECT categoryID, categoryName, categoryDescription FROM categories")
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}
	defer catRows.Close()

	// Collect the categories
	var categories []structs.Category
	for catRows.Next() {
		var c structs.Category
		if err := catRows.Scan(&c.CategoryID, &c.CategoryName, &c.CategoryDescription); err != nil {
			continue
		}
		categories = append(categories, c)
	}

	// Retrieve the users from the database
	userRows, err := db.DB.Query("SELECT userID, userUsername FROM users")
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}
	defer userRows.Close()

	// Collect the users
	var users []structs.User
	for userRows.Next() {
		var u structs.User
		if err := userRows.Scan(&u.UserID, &u.UserUsername); err != nil {
			continue
		}
		users = append(users, u)
	}

	// Render the template with posts, categories, and users
	RenderTemplate(w, "/", struct {
		IsAuthenticated bool
		Posts           []structs.Post
		Categories      []structs.Category
		Users           []structs.User
	}{
		IsAuthenticated: isAuthenticated,
		Posts:           posts,
		Categories:      categories,
		Users:           users,
	})
}
