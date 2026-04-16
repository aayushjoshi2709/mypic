package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct{
	RedisClient *redis.Client
}


var redisInstance *Redis

func Init() *Redis{
	if redisInstance == nil {
		return redisInstance
	}
	redisInstance = GetConn()
	return redisInstance
}

func GetConn() *Redis {
	redisURI = os.Getenv("REDIS_URI")
	redisPassword = os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisURI,
		Password: redisPassword,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	redisInstance = &Redis{
		RedisClient: client,
	}

	return redisInstance
}
