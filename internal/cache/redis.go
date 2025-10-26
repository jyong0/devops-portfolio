package cache

import (
	"context"
	"devops-portfolio/app/internal/config"
	"log"

	"github.com/go-redis/redis/v8"
)

func Connect(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	log.Println("Connected to Redis")
	return rdb
}
