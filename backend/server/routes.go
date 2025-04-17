package server

import (
    "Forum/backend/handler"
    "Forum/backend/middlewares"
    "Forum/backend/services"
    "net/http"

    "github.com/gorilla/mux"
)

func InitRoutes() {

    ////////////////////////////// Sub-router for guest //////////////////////////////////////////
    guestRouter := http.NewServeMux() // Create a new ServeMux for guest-related routes.
    guestRouter.HandleFunc("/", handler.GuestHome)                // Home page for guests
    guestRouter.HandleFunc("/hack", handler.CategoryHack)          // Category: Hack for guests
    guestRouter.HandleFunc("/prog", handler.CategoryProg)          // Category: Programming for guests
    guestRouter.HandleFunc("/news", handler.CategoryNews)          // Category: News for guests
    guestRouter.HandleFunc("/search", handler.SearchPseudo)       // Search functionality for guests
    guestRouter.HandleFunc("/about", handler.AboutForum)          // About the forum page

    // Apply middlewares (Logger and RateLimit) to the guestRouter for logging and rate limiting.
    protectedGuestRouter := middlewares.Logger(middlewares.RateLimit(guestRouter))

    ////////////////////////// Sub-router for authentication /////////////////////////////////////
    authRouter := http.NewServeMux() // Create a new ServeMux for authentication routes.
    authRouter.HandleFunc("/login", handler.Login)                // Login page
    authRouter.HandleFunc("/register", handler.Register)          // Registration page
    authRouter.HandleFunc("/google", handler.GoogleLogin)         // Google login
    authRouter.HandleFunc("/google/callback", handler.GoogleRegister) // Google registration callback

    // Apply middlewares (Logger and RateLimit) to the authRouter for logging and rate limiting.
    protectedAuthRouter := middlewares.Logger(middlewares.RateLimit(authRouter))

    ////////////////////////// Sub-router for user /////////////////////////////////////////////
    userRouter := http.NewServeMux() // Create a new ServeMux for user-related routes.
    userRouter.HandleFunc("/profile/{id}/edit", handler.UserEditProfile) // Edit user profile
    userRouter.HandleFunc("/posts/news", handler.UserCreatePost)         // Create a post in the "news" category
    userRouter.HandleFunc("/posts/hack", handler.UserCreatePost)         // Create a post in the "hack" category
    userRouter.HandleFunc("/posts/prog", handler.UserCreatePost)         // Create a post in the "programming" category
    userRouter.HandleFunc("/post/{id}/like", handler.ToggleLikePost)     // Like a post
    userRouter.HandleFunc("/post/{id}/comment", handler.ToggleLikeComment) // Like a comment on a post
    userRouter.HandleFunc("/logout", handler.Logout)                     // Log out user
    userRouter.HandleFunc("/profile", handler.UserProfile)               // View user profile

    // Apply middlewares (Logger, Authentication, and RateLimit) to userRouter.
    // Authentication middleware is applied to protect user-specific routes.
    protectedUserRouter := middlewares.Logger(middlewares.Authentication(middlewares.RateLimit(userRouter)))

    ////////////////////////// Sub-router for Admin ///////////////////////////////////////////
    adminRouter := mux.NewRouter() // Create a new router for admin-related routes.
    adminRouter.HandleFunc("/dashboard", handler.AdminDashboard)          // Admin dashboard
    adminRouter.HandleFunc("/user/{id}/delete", handler.AdminDeleteUser)   // Delete a user (admin only)
    adminRouter.HandleFunc("/comment/{id}/delete", handler.AdminDeleteComment) // Delete a comment (admin only)
    adminRouter.HandleFunc("/post/{id}/delete", handler.AdminDeletePost)   // Delete a post (admin only)

    // Apply middlewares (Logger, Authentication, and AdminAuthorization) to the adminRouter.
    // Authentication and AdminAuthorization middlewares are applied to protect admin-specific routes.
    protectedAdminRouter := middlewares.Logger(
        middlewares.Authentication(
            middlewares.AdminAuthorization(adminRouter),
        ),
    )

    ////////////////////////// Main router /////////////////////////////////////////////////
    mainRouter := http.NewServeMux() // Create a new ServeMux for the main router.

    // Attach the sub-routers to the main router with their respective URL prefixes.
    mainRouter.Handle("/forum/", http.StripPrefix("/forum", protectedGuestRouter))    // Guest routes (Forum)
    mainRouter.Handle("/auth/", http.StripPrefix("/auth", protectedAuthRouter))      // Authentication routes (Login/Register)
    mainRouter.Handle("/user/", http.StripPrefix("/user", protectedUserRouter))      // User-specific routes (Profile, Posts, etc.)
    mainRouter.Handle("/admin/", http.StripPrefix("/admin", protectedAdminRouter))  // Admin-specific routes (Dashboard, Delete Users, etc.)

    // Set the mainRouter as the global main router in the services structure for the Forum.
    services.F.MainRouter = mainRouter
}
