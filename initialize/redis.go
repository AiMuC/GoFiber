package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func RedisDrive() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("REDIS.DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS.ADDR"),
		Password: os.Getenv("REDIS.PASSWORD"),
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("无法连接到Redis服务器: %v", err))
	}
	return client
}
