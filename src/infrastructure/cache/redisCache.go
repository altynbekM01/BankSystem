package cache

import (
	"GolangwithFrame/src/domain/model"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/goccy/go-json"
	"time"
)

type RedisCache struct {
	host string
	db   int
	exp  time.Duration
}

func (cache *RedisCache) Set(key string, value *model.Product) {
	client := cache.getClient()
	ctx := context.Background()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(ctx, key, json, cache.exp*time.Second)
}

func (cache *RedisCache) Get(key string) *model.Product {
	client := cache.getClient()
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}
	product := model.Product{}
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		panic(err)
	}
	return &product

}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func NewRedisCache(host string, db int, exp time.Duration) ProductCache {
	return &RedisCache{
		host: host,
		db:   db,
		exp:  exp,
	}
}
