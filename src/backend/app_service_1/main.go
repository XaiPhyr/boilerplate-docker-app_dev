package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// go func() {
	// 	sub := client.Subscribe(ctx, "app_service")

	// 	ch := sub.Channel()
	// 	for msg := range ch {
	// 		log.Printf("APP SERVICE %s REDIS SUB: %s", service, msg.Payload)
	// 	}

	// 	defer sub.Close()
	// }()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/routed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CONNECTED TO APP SERVICE 1")
		w.Write([]byte("HELLO WORLD"))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Error ListenAndServe: %s", err)
}
