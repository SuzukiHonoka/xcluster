package session

import "xcluster/pkg/redis"

var store *Store

func InitStore(client *redis.Client) {
	store = NewStore(client)
}
