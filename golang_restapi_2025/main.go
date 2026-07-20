package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zhangguo2002/golangrestapi/dbconfig"
	"github.com/zhangguo2002/golangrestapi/internal/handlers"
	"github.com/zhangguo2002/golangrestapi/internal/routes"
	"github.com/zhangguo2002/golangrestapi/internal/store"
	"github.com/zhangguo2002/golangrestapi/serverconfig"
)

func main() {
	//Load config
	config, err := serverconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
	//connect to db
	db := dbconfig.ConnectDB(config.DatabaseURL)
	defer db.Close()
	dbconfig.RunMigrations(db, "internal/migrations/schema.sql")
	queries := store.New(db)
	//Create a new handler
	handler := handlers.NewHandlers(db, queries)
	//set up the HTTP server
	mux := http.NewServeMux()
	//Setup Routes
	routes.SetupRoutes(mux, handler)
	//server instance
	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}
	fmt.Printf("Server is up and running on PORT %s\n", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed %v", err)
	}
}
