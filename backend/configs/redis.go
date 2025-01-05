package configs

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis(config *Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	return rdb, nil
}
