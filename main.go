package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	client := createClient()

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = writeName(client, "Elliot")
	if err != nil {
		log.Fatal(err)
	}

	val, err := readName(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}

func createClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func writeName(client *redis.Client, name string) error {
	return client.Set(context.Background(), "name", name, 0).Err()
}

func readName(client *redis.Client) (string, error) {
	return client.Get(context.Background(), "name").Result()
}
