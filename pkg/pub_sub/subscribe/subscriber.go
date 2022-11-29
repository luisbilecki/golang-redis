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

	subscription := client.Subscribe(ctx, "mychannel")
	defer subscription.Close()
	for {
		msg, err := subscription.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}
}
