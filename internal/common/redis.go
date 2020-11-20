package common

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	fmt.Println("redisClient init")
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
