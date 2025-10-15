package store

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const CacheDuration = 24 * time.Hour

var (
	ctx          = context.Background()
	storeService = StoreService{}
)

type StoreService struct {
	redisClient *redis.Client
}

func InitstoreService() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("failed to connect to redis %v", err)
	}
	log.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return &storeService
}

func SaveUrlMapping(shorturl, longurl, userid string) {
	if err := storeService.redisClient.Set(ctx, shorturl, longurl, CacheDuration).Err(); err != nil {
		log.Printf("failed to save url mapping %v", err)
	}
}

func RetrieveInitialUrl(shorturl string) string {
	Result, err := storeService.redisClient.Get(ctx, shorturl).Result()
	if err != nil {
		log.Printf("failed to retrieve initial url %v", err)
		return ""
	}
	return Result
}
