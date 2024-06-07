package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "",
	DB:       0,
})

func main() {
	r := chi.NewRouter()

	port := os.Getenv("HTTP_PORT")
	version := os.Getenv("VERSION")
	service := os.Getenv("SERVICE")

	log.SetFlags(log.Llongfile | log.LstdFlags)

	fmt.Println()
	log.Printf("-> Service: %s", service)
	log.Printf("-> Version: %s", version)
	log.Printf("-> Local:   http://localhost:%s", port)

	go func() {
		sub := client.Subscribe(ctx, "app_service")

		ch := sub.Channel()
		for msg := range ch {
			log.Printf("APP SERVICE %s REDIS SUB: %s", service, msg.Payload)
		}

		defer sub.Close()
	}()

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Error ListenAndServe: %s", err)
}
