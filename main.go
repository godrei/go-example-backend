package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	client := createClient()

	pong, err := pingDatabase(client)
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

func pingDatabase(client *redis.Client) (string, error) {
	return client.Ping(context.Background()).Result()
}

func writeName(client *redis.Client, name string) error {
	return client.Set(context.Background(), "name", name, 0).Err()
}

func readName(client *redis.Client) (string, error) {
	return client.Get(context.Background(), "name").Result()
}
