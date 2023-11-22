package session

import (
	"time"
	"xcluster/pkg/redis"
)

type Store struct {
	RedisClient *redis.Client
}

func NewStore(client *redis.Client) *Store {
	return &Store{RedisClient: client}
}

func (s *Store) Add(session *Session) error {
	duration := session.Lease.ExpirationTime.Sub(time.Now())
	if _, err := s.RedisClient.Set(string(session.ID), session.Lease, duration).Result(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Get(id ID) (*Session, error) {
	var lease Lease
	if err := s.RedisClient.Get(string(id)).Scan(&lease); err != nil {
		return nil, err
	}
	session := &Session{
		ID:    id,
		Lease: &lease,
	}
	return session, nil
}

func (s *Store) Delete(id ID) error {
	_, err := s.RedisClient.Del(string(id)).Result()
	return err
}
