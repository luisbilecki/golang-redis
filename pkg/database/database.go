package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

const (
	REDIS_URL      = "localhost:6379"
	REDIS_PASSWORD = "" // No password set
	REDIS_DB       = 0  // Database default
)

type Character struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	character := Character{ID: uuid.New(), Name: "Eivor"}
	characterAsJson, _ := json.Marshal(character)
	key := character.ID.String()
	err := client.Set(ctx, key, characterAsJson, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	recoveredValue, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Key %s has value %s\n", key, recoveredValue)

	var newCharacter Character
	_ = json.Unmarshal([]byte(recoveredValue), &newCharacter)
	fmt.Println(newCharacter)
}
