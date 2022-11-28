package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

const (
	REDIS_URL      = "localhost:6379"
	REDIS_PASSWORD = "" // No password set
	REDIS_DB       = 0  // Database default
)

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Printf("Pong Response: %v, Error: %v\n", pong, err)
}
