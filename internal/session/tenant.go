package session

import (
	"encoding/json"
	"time"
	"xcluster/internal/user"
)

type Lease struct {
	TenantID   user.ID
	ExpireTime time.Time
}

func NewLease(id user.ID, duration time.Duration) *Lease {
	return &Lease{
		TenantID:   id,
		ExpireTime: time.Now().Add(duration),
	}
}

func (l *Lease) Expire() bool {
	now := time.Now()
	return now.Sub(l.ExpireTime) > 0
}

// MarshalBinary is an implementation of encoding.BinaryMarshaler
func (l *Lease) MarshalBinary() (data []byte, err error) {
	return json.Marshal(*l)
}

// UnmarshalBinary is an implementation of encoding.BinaryUnmarshaler
func (l *Lease) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
