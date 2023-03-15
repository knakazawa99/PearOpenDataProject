package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"api/domain/repository"
)

type cache struct {
}

func (c cache) Set(key string, value string, ttl time.Duration) error {
	ctx := context.Background()
	client := connectRedis()
	if err := client.Set(ctx, key, value, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func (c cache) Get(key string) (string, error) {
	ctx := context.Background()
	client := connectRedis()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val, nil
}

func connectRedis() *redis.Client {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("failed: load env value : %v", err)
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       0,
	})
	return rdb
}

func NewCache() repository.Cache {
	return &cache{}
}
