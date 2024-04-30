package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	client := createDBClient()

	pong, err := pingDB(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	err = writeNameToDB(client, "Elliot")
	if err != nil {
		log.Fatal(err)
	}

	val, err := readNameFromDB(client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}

func createDBClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     dbAddress(),
		Password: "",
		DB:       0,
	})
}

func pingDB(client *redis.Client) (string, error) {
	return client.Ping(context.Background()).Result()
}

func writeNameToDB(client *redis.Client, name string) error {
	return client.Set(context.Background(), "name", name, 0).Err()
}

func readNameFromDB(client *redis.Client) (string, error) {
	return client.Get(context.Background(), "name").Result()
}

func dbAddress() string {
	addr := os.Getenv("REDIS_DSN")
	if addr != "" {
		return addr
	}
	return "localhost:6379"
}
