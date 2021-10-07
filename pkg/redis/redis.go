package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	//TODO: use ENV VAR
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
