package redis

import (
	"gateway/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Rdb *redis.Client
}

func New(cfg *config.Config) *Client {
	return &Client{Rdb: redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPw,
		DB:       0,
	})}
}
