package redis

import (
	"context"
	"file/internal/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis(cfg *config.RedisConfig) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	return err
}

func Close() {
	if RedisClient != nil {
		RedisClient.Close()
	}
}
