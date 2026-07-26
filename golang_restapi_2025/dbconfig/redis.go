package dbconfig

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func ConnectRedis(redisAddr string, redisPassword string) *redis.Client {
	db := 0

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       db,
	})

	//test connection
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis:%v", err))
	}
	fmt.Println("Connected to redis successfully")
	return rdb
}
