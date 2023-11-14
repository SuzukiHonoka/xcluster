package session

import (
	"fmt"
	"github.com/google/uuid"
)

type Session struct {
	ID    ID
	Lease *Lease
}

func init() {
	// use rand pool by default for stat
	uuid.EnableRandPool()
}

func NewSession(lease *Lease) (*Session, error) {
	// generate uuid
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	session := &Session{
		ID:    ID(uid.String()),
		Lease: lease,
	}
	// save to redis
	if err = store.Add(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s Session) String() string {
	return fmt.Sprintf("[session] id=%s, tenantID=%d, expire=%s", s.ID, s.Lease.TenantID, s.Lease.ExpireTime)
}

func (s Session) Delete() error {
	return s.ID.DeleteSession()
}
