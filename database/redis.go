package database

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

// RDB ...
var RDB *redis.Client

// BuildRedisConnection ...
func NewRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "6389",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if RDB == nil {
		log.Fatal("連接 redis 出現錯誤")
	}

	fmt.Println("Redis連線成功...接著是Casbin...")
}
