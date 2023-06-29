package initialize

import (
	"context"
	"fmt"
	"github.com/aimuc/gofiber/support"
	"github.com/redis/go-redis/v9"
)

func RedisDrive() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     support.Env("REDIS.ADDR", "127.0.0.1:6379").(string),
		Password: support.Env("REDIS.PASSWORD", "").(string),
		DB:       support.Env("REDIS.DB", 0).(int),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("无法连接到Redis服务器: %v", err))
	}
	return client
}
