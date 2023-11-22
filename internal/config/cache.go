package config

import (
	"errors"
	"sync"
)

var cache *Cache

type Cache struct {
	store *sync.Map
}

func NewCache() *Cache {
	return &Cache{
		store: &sync.Map{},
	}
}

func (c *Cache) Get(key string) (string, error) {
	val, ok := c.store.Load(key)
	if !ok {
		return "", errors.New("key not found")
	}

	return val.(string), nil
}

func (c *Cache) Set(key string, val string) {
	c.store.Store(key, val)
}
