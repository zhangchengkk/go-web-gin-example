package config

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var RedisContext context.Context

func initRedis() {
	RedisContext = context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "dachuizi", // no password set
		DB:       0,  // use default DB
	})
}
