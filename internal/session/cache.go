package session

import (
	"context"
	"xcluster/pkg/redis"
)

var store Store

func InitStore(client *redis.Client) {
	// todo: use specific store according to configuration
	store = NewRedisStore(context.Background(), client)
}
