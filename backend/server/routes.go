package server

import (
	"Forum/backend/handler"
	"Forum/backend/middlewares"
	"Forum/backend/structs"
	"net/http"
	
)

var F =  &structs.Forum{}

func InitRoutes()  {
//////////////////////////////sous router pour guest //////////////////////////////////////////
	guestRouter := http.NewServeMux()
	guestRouter.HandleFunc("/", handler.GuestHome)
	guestRouter.HandleFunc("/hack", handler.GuestHack)
	guestRouter.HandleFunc("/prog", handler.GuestProg)
	guestRouter.HandleFunc("/news", handler.GuestNews)
	guestRouter.HandleFunc("/search", handler.GuestSearch)
	guestRouter.HandleFunc("/about", handler.GuestAbout)

	protectedGuestRouter := middlewares.Logger(middlewares.RateLimit(guestRouter))
////////////////////////sous-router pour authentification/////////////////////////////////////
	authRouter := http.NewServeMux()
	authRouter.HandleFunc("/login", handler.Login)
	authRouter.HandleFunc("/register",  handler.Register)

	protectedAuthRouter := middlewares.Logger(middlewares.RateLimit(authRouter))
/////////////////////////sous router pour user ////////////////////////////////////////
	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/profile/{id}/edit", handler.UserEditProfile)
	userRouter.HandleFunc("/posts/news", handler.UserCreatePost)
	userRouter.HandleFunc("/posts/hack", handler.UserCreatePost)
	userRouter.HandleFunc("/posts/prog", handler.UserCreatePost)
	userRouter.HandleFunc("/post/{id}/like", handler.UserLikePost) 
	userRouter.HandleFunc("/post/{id}/comment", handler.UserAddComment) 
	userRouter.HandleFunc("/logout", handler.Logout)
	userRouter.HandleFunc("/profile", handler.UserProfile)

	protectedUserRouter := middlewares.Logger(middlewares.Authentication(middlewares.RateLimit(userRouter)))    //(protégées par le middlewares Authentication


///////////////////////// sous routeur mour Admin ////////////////////////////////////////////////
	adminRouter := http.NewServeMux()  
	adminRouter.HandleFunc("/dashboard", handler.AdminDashboard)
	adminRouter.HandleFunc("/user/{id}/delete", handler.AdminDeleteUser) 
	adminRouter.HandleFunc("/comment/{id}/delete", handler.AdminDeleteComment) 
	adminRouter.HandleFunc("/post/{id}/delete", handler.AdminDeletePost) 

	protectedAdminRouter := middlewares.Logger(middlewares.Authentication(middlewares.AdminAuthorization(adminRouter))) //(protégées par les middelwares AdminAuthorization + Authentication
///////////////////////// routeur Principal ////////////////////////////////////////////////

	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/overload", handler.OverloadHandler)
	mainRouter.Handle("/forum/", http.StripPrefix("/forum", protectedGuestRouter))
	mainRouter.Handle("/auth/", http.StripPrefix("/auth", protectedAuthRouter))
	mainRouter.Handle("/user/", http.StripPrefix("/user", protectedUserRouter))
	mainRouter.Handle("/admin/", http.StripPrefix("/admin", protectedAdminRouter))

	F.MainRouter = mainRouter       /////////ajout du routeur principal a la structure globlal Forum 


}
