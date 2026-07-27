package routes

import (
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/handlers"
	"github.com/zhangguo2002/golangrestapi/internal/middlewares"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	userMux := http.NewServeMux()
	userMux.Handle("GET /all", http.HandlerFunc(handler.ListUserHandler()))
	userMux.HandleFunc("POST /register", handler.CreateUserHandler())
	userMux.HandleFunc("POST /login", handler.LoginUserHandler())
	userMux.Handle("GET /profile", middlewares.AuthMiddleware(http.HandlerFunc(handler.UserProfile())))
	userMux.Handle("POST /session/logout", middlewares.AuthMiddleware(http.HandlerFunc(handler.LogoutHandler())))
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

	//configure the upload mux
	uploadMux := http.NewServeMux()
	uploadMux.Handle("POST /", middlewares.AuthMiddleware(http.HandlerFunc(handler.UploadProfileImageHandler())))
	mux.Handle("/upload/", http.StripPrefix("/upload", uploadMux))
}
