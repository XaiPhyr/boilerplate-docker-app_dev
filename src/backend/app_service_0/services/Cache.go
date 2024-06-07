package services

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Cache struct{}

var message chan string
var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "", // no password set
	DB:       0,  // use default DB
})

func (c Cache) SetCache(data interface{}) {
	err := client.Set(ctx, "foo", data, 0).Err()
	if err != nil {
		log.Printf("Error Set Cache: %s", err)
		return
	}
}

func (c Cache) GetCache() {
	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		log.Printf("Error Get Cache: %s", err)
		return
	}

	log.Printf("Get Cache: %s", val)
}

func (c Cache) RedisPub(data string) {
	pub := client.Publish(ctx, "app_service", data)
	if pub != nil {
		log.Printf("Redis Pub: %s", pub)
		return
	}
}

func (c Cache) RedisSub() {
	sub := client.Subscribe(ctx, "app_service")
	defer sub.Close()

	ch := sub.Channel()
	for msg := range ch {
		message <- msg.Payload
	}
}
