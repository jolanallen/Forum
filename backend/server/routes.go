package server

import (
    "Forum/backend/handler"
    "Forum/backend/middlewares"
    "Forum/backend/services"
    "net/http"

    "github.com/gorilla/mux"
)

func InitRoutes() {
    fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    //////////////////////////////sous router pour guest //////////////////////////////////////////
    guestRouter := http.NewServeMux()
    guestRouter.HandleFunc("/", handler.GuestHome)
    guestRouter.HandleFunc("/hack", handler.CategoryHack)
    guestRouter.HandleFunc("/prog", handler.CategoryProg)
    guestRouter.HandleFunc("/news", handler.CategoryNews)
    guestRouter.HandleFunc("/search", handler.SearchPseudo)
    guestRouter.HandleFunc("/about", handler.AboutForum)

    protectedGuestRouter := middlewares.Logger(middlewares.RateLimit(guestRouter))
    ////////////////////////sous-router pour authentification/////////////////////////////////////
    authRouter := http.NewServeMux()
    authRouter.HandleFunc("/login", handler.Login)
    authRouter.HandleFunc("/register", handler.Register)
    authRouter.HandleFunc("/google", handler.GoogleLogin)
    authRouter.HandleFunc("/google/callback", handler.GoogleRegister)
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

    protectedUserRouter := middlewares.Logger(middlewares.Authentication(middlewares.RateLimit(userRouter))) //(protégées par le middlewares Authentication

    ///////////////////////// sous routeur mour Admin ////////////////////////////////////////////////
    adminRouter := mux.NewRouter()
    adminRouter.HandleFunc("/dashboard", handler.AdminDashboard)
    adminRouter.HandleFunc("/user/{id}/delete", handler.AdminDeleteUser)
    adminRouter.HandleFunc("/comment/{id}/delete", handler.AdminDeleteComment)
    adminRouter.HandleFunc("/post/{id}/delete", handler.AdminDeletePost)

    protectedAdminRouter := middlewares.Logger(
        middlewares.Authentication(
            middlewares.AdminAuthorization(adminRouter),
        ),
    ) //(protégées par les middelwares AdminAuthorization + Authentication
    ///////////////////////// routeur Principal ////////////////////////////////////////////////

    mainRouter := http.NewServeMux()
    mainRouter.Handle("/forum/", http.StripPrefix("/forum", protectedGuestRouter))
    mainRouter.Handle("/auth/", http.StripPrefix("/auth", protectedAuthRouter))
    mainRouter.Handle("/user/", http.StripPrefix("/user", protectedUserRouter))
    mainRouter.Handle("/admin/", http.StripPrefix("/admin", protectedAdminRouter))
        
    services.F.MainRouter = mainRouter /////////ajout du routeur principal a la structure globlal Forum

}
