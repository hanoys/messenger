package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(host string, port string) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr: host + ":" + port,
    })
}
