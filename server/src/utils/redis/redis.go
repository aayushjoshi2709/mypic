package redis

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	RedisClient *redis.Client
}

var redisInstance *Redis

func Init() *Redis {
	if redisInstance != nil {
		return redisInstance
	}
	redisInstance = &Redis{
		RedisClient: GetConn(),
	}
	return redisInstance
}

func GetConn() *redis.Client {

	rediUri := os.Getenv("REDIS_URI")
	if rediUri == "" {
		panic("REDIS_URI environment variable not set")
	}

	opt, err := redis.ParseURL(rediUri)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	client.Ping(context.Background())

	return client
}

func (r *Redis) Set(ctx *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return r.RedisClient.Set(ctx, key, value, expiration).Err()
}

func (r *Redis) Get(ctx *gin.Context, key string) (string, error) {
	return r.RedisClient.Get(ctx, key).Result()
}

func (r *Redis) GetAndDelete(ctx *gin.Context, key string) (string, error) {
	luaScript := `
		local value = redis.call("GET", KEYS[1])
		if value then
			redis.call("DEL", KEYS[1])
		end
		return value
	`
	result, err := r.RedisClient.Eval(ctx, luaScript, []string{key}).Result()
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", redis.Nil
	}
	return result.(string), nil
}

func (r *Redis) Del(ctx *gin.Context, keys ...string) error {
	return r.RedisClient.Del(ctx, keys...).Err()
}

func (r *Redis) Exists(ctx *gin.Context, key string) (bool, error) {
	result, err := r.RedisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *Redis) BulkSet(ctx *gin.Context, keyValuePairs map[string]string, expiration time.Duration) error {
	pipe := r.RedisClient.Pipeline()
	for key, value := range keyValuePairs {
		pipe.Set(ctx, key, value, expiration)
	}
	_, err := pipe.Exec(ctx)
	return err
}
