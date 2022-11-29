package main

import (
	"context"
	"fmt"
	"time"

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

	key := "CarBrands"
	value := "Acura,Alfa Romeo,Audi,BMW,Chevrolet"
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	recoveredValue, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Key %s has value %s\n", key, recoveredValue)

	// Expires in 5 seconds
	client.Set(ctx, "ExpiresIn5Sec", "Will Expire Soon", 5*time.Second)

}
