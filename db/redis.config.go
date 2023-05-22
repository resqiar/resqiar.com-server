package db

import (
	"os"

	redis "github.com/gofiber/storage/redis/v2"
)

var RedisStore *redis.Storage

func InitRedis() {
	REDIS_URL := os.Getenv("REDIS_URL")

	RedisStore = redis.New(redis.Config{
		URL:   REDIS_URL,
		Reset: false,
	})
}
