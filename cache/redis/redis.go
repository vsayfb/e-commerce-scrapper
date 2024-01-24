package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/vsayfb/e-commerce-scrapper/product"
)

type RedisCache struct {
	client *redis.Client
}

func New() *RedisCache {
	url := "redis://localhost:6379/"

	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rc := redis.NewClient(opts)

	ctx := context.Background()

	if err := rc.Ping(ctx).Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to Redis.")
	}

	return &RedisCache{
		client: rc,
	}
}

func (r *RedisCache) Add(key string, data interface{}) bool {
	ctx := context.Background()

	err := r.client.Set(ctx, key, data, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("Data is cached successfully.")
		return true
	}
}

func (r *RedisCache) GetProducts(key string) ([]product.Product, error) {
	ctx := context.Background()

	res, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, errors.New(fmt.Sprintf("Key - %v - does not exist \n", key))
	} else if err != nil {
		return nil, errors.New(err.Error())
	}

	var products []product.Product

	encodingErr := json.Unmarshal([]byte(res), &products)

	if err != nil {
		return nil, errors.New(encodingErr.Error())
	}

	return products, nil
}
