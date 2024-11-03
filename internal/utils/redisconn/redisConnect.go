package redisconn

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient is a global Redis client instance
var RedisClient *redis.Client

func ConnectRedis(cfg *RedisConfig) *redis.Client {
	address := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})
	// Test the connection to Redis
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Ping to check if the connection is successful
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Failed to connect to Redis at %s - %v", address, err)
	} else {
		log.Println("Connected to Redis successfully")
	}
	return RedisClient
}
