package configs

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	RedisClient          *redis.Client
	MaxRequestsPerSecond int
	BlockDuration        time.Duration
}

func LoadConfig() (*Config, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &Config{
		RedisClient:          redisClient,
		MaxRequestsPerSecond: 5,
		BlockDuration:        2 * time.Second,
	}, nil
}
