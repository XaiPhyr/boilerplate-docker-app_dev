package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	sql "app_backend/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

var endpoint = os.Getenv("ENDPOINT")
var version = os.Getenv("VERSION")
var url = func(api string) string {
	return endpoint + api
}

func main() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "app_redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	port := os.Getenv("HTTP_PORT")
	log.SetFlags(log.Llongfile | log.LstdFlags)
	sql.MigrateSQL()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.GetHead)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Get(endpoint, func(w http.ResponseWriter, r *http.Request) {
		err := client.Set(ctx, "foo", "SET REDIS", 0).Err()
		if err != nil {
			panic(err)
		}

		w.Write([]byte("WELCOME!"))
	})

	r.Route(url("/hello"), func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("HELLO WORLD"))
		})
	})

	r.Route(url("/get-redis"), func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			val, err := client.Get(ctx, "foo").Result()
			if err != nil {
				panic(err)
			}

			w.Write([]byte(fmt.Sprintf("Redis Get: %s", val)))
		})
	})

	fmt.Println()
	log.Printf("-> Version: %s", version)
	log.Printf("-> Local:   http://localhost:8200")

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
