package session

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// Prefix is a common Prefix, for now
const Prefix = "session-"

type Session struct {
	ID    ID
	Lease *Lease
}

func NewSession(lease *Lease) (*Session, error) {
	// generate uuid
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	session := &Session{
		ID:    ID(Prefix + uid.String()),
		Lease: lease,
	}
	// save to redis
	if err = store.Add(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s Session) String() string {
	return fmt.Sprintf("[session] id=%s, tenantID=%d, expire=%s", s.ID, s.Lease.TenantID, s.Lease.ExpirationTime)
}

func (s Session) PrettyString() string {
	return fmt.Sprintf("session: id=%s (tenantID=%d, expire=%s)",
		s.ID, s.Lease.TenantID, s.Lease.ExpirationTime.Format(time.DateTime))
}

func (s Session) ShortString() string {
	return fmt.Sprintf("session: id=" + string(s.ID))
}

func (s Session) Delete() error {
	return s.ID.DeleteSession()
}
