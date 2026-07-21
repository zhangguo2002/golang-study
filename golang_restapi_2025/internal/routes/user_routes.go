package routes

import (
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/handlers"
)

func SetupUserRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	userMux := http.NewServeMux()
	userMux.HandleFunc("POST /register", handler.CreateUserHandler())
	userMux.HandleFunc("POST /login", handler.LoginUserHandler())
	mux.Handle("/users/", http.StripPrefix("/users", userMux))
}
