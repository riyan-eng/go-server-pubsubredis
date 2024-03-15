package infrastructure

import (
	"context"
	"fmt"
	"os"

	"server/env"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%v:%v", env.NewEnvironment().REDIS_HOST, env.NewEnvironment().REDIS_PORT)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: env.NewEnvironment().REDIS_USERNAME,
		Password: env.NewEnvironment().REDIS_PASSWORD,
		DB:       env.NewEnvironment().REDIS_DATABASE,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		fmt.Printf("redis: can't ping to redis - %v \n", err)
		os.Exit(1)
	}
	fmt.Println("redis: connection opened")
}
