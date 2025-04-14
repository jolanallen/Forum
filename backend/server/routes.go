package server

import (
	"Forum/backend/handler"
	"Forum/backend/middlewares"
	"Forum/backend/services"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {

	//////////////////////////////sous router pour guest //////////////////////////////////////////
	guestRouter := http.NewServeMux()
	guestRouter.HandleFunc("/", handler.GuestHome)
	guestRouter.HandleFunc("/hack", handler.CategoryHack)
	guestRouter.HandleFunc("/prog", handler.CategoryProg)
	guestRouter.HandleFunc("/news", handler.CategoryNews)
	guestRouter.HandleFunc("/search", handler.SearchPseudo)
	guestRouter.HandleFunc("/about", handler.AboutForum)

	// Applique le logger et limite les requêtes pour les routes guest
	protectedGuestRouter := middlewares.Logger(middlewares.RateLimit(guestRouter))

	////////////////////////sous-router pour authentification/////////////////////////////////////
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("/login", handler.Login)
	authRouter.HandleFunc("/register", handler.Register)
	// Routes d'authentification ajoutées ici
	authRouter.HandleFunc("/google", handler.GoogleLogin)
	authRouter.HandleFunc("/google/callback", handler.GoogleRegister)

	// Applique le logger et limite les requêtes pour les routes d'authentification
	protectedAuthRouter := middlewares.Logger(middlewares.RateLimit(authRouter))

	/////////////////////////sous router pour user ////////////////////////////////////////
	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/profile/{id}/edit", handler.UserEditProfile)
	userRouter.HandleFunc("/posts/news", handler.UserCreatePost)
	userRouter.HandleFunc("/posts/hack", handler.UserCreatePost)
	userRouter.HandleFunc("/posts/prog", handler.UserCreatePost)
	userRouter.HandleFunc("/post/{id}/like", handler.ToggleLikePost)
	userRouter.HandleFunc("/post/{id}/comment", handler.ToggleLikeComment)
	userRouter.HandleFunc("/logout", handler.Logout)
	userRouter.HandleFunc("/profile", handler.UserProfile)

	// Applique le middleware d'authentification + rate limiting pour les routes utilisateurs
	protectedUserRouter := middlewares.Logger(middlewares.Authentication(middlewares.RateLimit(userRouter)))

	///////////////////////// sous routeur pour Admin ////////////////////////////////////////////////
	adminRouter := mux.NewRouter()
	adminRouter.HandleFunc("/dashboard", handler.AdminDashboard)
	adminRouter.HandleFunc("/user/{id}/delete", handler.AdminDeleteUser)
	adminRouter.HandleFunc("/comment/{id}/delete", handler.AdminDeleteComment)
	adminRouter.HandleFunc("/post/{id}/delete", handler.AdminDeletePost)

	// Applique le middleware d'authentification + autorisation admin + rate limiting pour les routes admin
	protectedAdminRouter := middlewares.Logger(
		middlewares.Authentication(
			middlewares.AdminAuthorization(adminRouter),
		),
	)

	///////////////////////// routeur Principal ////////////////////////////////////////////////

	// Création du routeur principal
	mainRouter := http.NewServeMux()
	// Redirige les routes de forum vers celles de guestRouter
	mainRouter.Handle("/forum/", http.StripPrefix("/forum", protectedGuestRouter))
	// Redirige les routes d'authentification vers celles de authRouter
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", protectedAuthRouter))
	// Redirige les routes d'utilisateur vers celles de userRouter
	mainRouter.Handle("/user/", http.StripPrefix("/user", protectedUserRouter))
	// Redirige les routes d'administration vers celles de adminRouter
	mainRouter.Handle("/admin/", http.StripPrefix("/admin", protectedAdminRouter))
		
	// Enregistre le routeur principal dans la structure globale du forum
	services.F.MainRouter = mainRouter
}
