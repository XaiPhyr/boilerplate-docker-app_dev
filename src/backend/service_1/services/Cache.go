package services

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type (
	Cache struct{}

	RedisPubSub struct {
		User    string `json:"user"`
		Message string `json:"message"`
	}
)

var message chan string
var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func (c Cache) SetCache(key, value string) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Printf("Error Set Cache: %s", err)
		return
	}
}

func (c Cache) GetCache(key string) string {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Printf("Error Get Cache: %s", err)
		return ""
	}

	log.Printf("Get Cache: %s", val)
	return val
}

func (c Cache) RedisPub(data string) {
	pub := client.Publish(ctx, "api_service_1", data)
	if pub != nil {
		log.Printf("Redis Pub: %s", pub)
		return
	}
}

func (c Cache) RedisSub() {
	sub := client.Subscribe(ctx, "api_service_1")
	defer sub.Close()

	ch := sub.Channel()
	for msg := range ch {
		message <- msg.Payload
	}
}
