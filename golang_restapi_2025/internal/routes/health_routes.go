package routes

import (
	"net/http"

	"github.com/zhangguo2002/golangrestapi/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux,handler *handlers.Handler){
	mux.HandleFunc("/health",handler.HealthHandler())
}