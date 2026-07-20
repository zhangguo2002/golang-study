package routes

import (
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/handlers"
)

func SetupUserRoutes(mux *http.ServeMux,handler *handlers.Handler){
	mux.HandleFunc("POST /user/register",handler.CreateUserHandler())
}