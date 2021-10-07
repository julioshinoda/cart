package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/julioshinoda/cart/entity"
	"github.com/julioshinoda/cart/internal/usecase/cart"
	rds "github.com/julioshinoda/cart/pkg/redis"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewCartRedis(client *redis.Client) cart.Repository {
	return &RedisClient{
		client: rds.NewRedisClient(),
	}
}

func (rc *RedisClient) Get(id string) (entity.Cart, error) {
	val, err := rc.client.Get(ctx, id).Result()
	if err != nil {
		return entity.Cart{}, err
	}
	var c entity.Cart
	if err := json.Unmarshal([]byte(val), &c); err != nil {
		return entity.Cart{}, err
	}
	return c, nil
}
func (rc *RedisClient) Update(e entity.Cart) error {
	return nil
}
