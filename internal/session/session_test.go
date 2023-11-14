package session

import (
	"errors"
	"math/rand"
	"testing"
	"time"
	"xcluster/pkg/redis"
)

var (
	addr     = "172.19.80.201:6379"
	password = ""
	n        = 10
	ttlMax   = 10 // in minutes since seconds might be too quick
	sessions = make([]*Session, 0, n)
)

func init() {
	config := redis.NewConfig(addr, password)
	client := redis.NewRedis(config)
	InitStore(client)
}

func generateSessions(n int, hook func(session *Session)) {
	for i := 0; i < n; i++ {
		ttl := time.Duration(rand.Intn(ttlMax-1)+1) * time.Minute
		session, err := NewSession(NewLease(0, ttl))
		if err != nil {
			panic(err)
		}
		hook(session)
	}
}

func TestNewSession(t *testing.T) {
	generateSessions(n, func(session *Session) {
		sessions = append(sessions, session)
		t.Log(session.String())
	})
}

func TestNewSessionFromUUID(t *testing.T) {
	for _, session := range sessions {
		s, err := session.ID.GetSession()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("redis: %s", s)
		if s.ID != session.ID {
			t.Error(errors.New("session uuid not identical"))
		}
	}
}

func TestDeleteSessionFromUUID(t *testing.T) {
	for _, session := range sessions {
		if err := session.ID.DeleteSession(); err != nil {
			t.Error(err)
		}
	}
}
