package session

import (
	"golang.org/x/net/context"
	"time"
	"xcluster/pkg/redis"
)

type RedisStore struct {
	Context     context.Context
	RedisClient *redis.Client
}

func NewRedisStore(context context.Context, client *redis.Client) Store {
	return &RedisStore{
		Context:     context,
		RedisClient: client,
	}
}

func (s *RedisStore) Add(session *Session) error {
	duration := session.Lease.ExpirationTime.Sub(time.Now())
	if _, err := s.RedisClient.Set(s.Context, string(session.ID), session.Lease, duration).Result(); err != nil {
		return err
	}
	return nil
}

func (s *RedisStore) Get(id ID) (*Session, error) {
	var lease Lease
	if err := s.RedisClient.Get(s.Context, string(id)).Scan(&lease); err != nil {
		return nil, err
	}
	session := &Session{
		ID:    id,
		Lease: &lease,
	}
	return session, nil
}

func (s *RedisStore) Delete(id ID) error {
	_, err := s.RedisClient.Del(s.Context, string(id)).Result()
	return err
}
