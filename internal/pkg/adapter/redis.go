package adapter

import (
	"context"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var redisOnce sync.Once
var redisClient *redis.Client

func GetRedisClient() (*redis.Client, error) {
	var err error
	redisOnce.Do(func() {
		redisClient, err = ConnectRedis(os.Getenv("REDIS_HOST"))
	})

	return redisClient, err
}

func ConnectRedis(host string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
