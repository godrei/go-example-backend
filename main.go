package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = client.Set(context.Background(), "name", "Elliot", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}
