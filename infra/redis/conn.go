package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/nemonicgod/terraforms-api/config"
)

// Connect ...
func Connect(c config.Reader) *redis.Client {
	redis_host := c.GetString(config.RedisHost)
	redis_port := c.GetString(config.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redis_host, redis_port),
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}
