package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Config *Config
	*redis.Client
}

func NewRedisWrapper(config *Config) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	// test client connection
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("redis: ping failed, err=%s", err)
	}
	return &Client{
		Config: config,
		Client: client,
	}, nil
}

func (c *Client) Close() error {
	return c.Close()
}
