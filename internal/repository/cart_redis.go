package repository

import (
	"context"
	"encoding/json"
	"errors"

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

func (rc *RedisClient) Get(id string) (cart *entity.Cart, err error) {
	val, err := rc.client.Get(ctx, id).Result()
	if errors.Is(err, redis.Nil) {
		return cart, entity.ErrNotFound
	}

	if err != nil {
		return cart, err
	}
	if err := json.Unmarshal([]byte(val), &cart); err != nil {
		return cart, err
	}
	return cart, nil
}
func (rc *RedisClient) Update(e *entity.Cart) error {
	toSave, err := json.Marshal(e)
	if err != nil {
		return err
	}
	return rc.client.Set(ctx, e.ID, toSave, -1).Err()
}
