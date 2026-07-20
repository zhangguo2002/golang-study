package dbconfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	fmt.Println("Connected to db")
	return db
}

func RunMigrations(db *sql.DB, schemaPath string) {
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}
	if _, err = db.Exec(string(schema)); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	fmt.Println("Migrations ran successfully")
}
