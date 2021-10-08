package redis

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	//TODO: use ENV VAR
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
