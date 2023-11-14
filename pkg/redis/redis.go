package redis

import (
	"github.com/go-redis/redis"
	"log"
)

type Client struct {
	Config *Config
	*redis.Client
}

func NewRedis(config *Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// test client
	if _, err := client.Ping().Result(); err != nil {
		log.Fatalf("redis: ping failed, err=%s", err)
	}
	return &Client{
		Config: config,
		Client: client,
	}
}
