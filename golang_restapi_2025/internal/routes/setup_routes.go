package routes

import (
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux,handler *handlers.Handler){
	SetupHealthRoute(mux,handler)
	SetupUserRoutes(mux,handler)
}