package handlers

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/zhangguo2002/golangrestapi/internal/store"
)

type Handler struct {
	//DB instance
	DB *sql.DB
	//Query stores
	Queries *store.Queries
	Redis   *redis.Client
}

func NewHandlers(db *sql.DB, queries *store.Queries, redisClient *redis.Client) *Handler {
	return &Handler{
		DB:      db,
		Queries: queries,
		Redis:   redisClient,
	}
}
