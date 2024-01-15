package session

import (
	"encoding/json"
	"time"
)

type Lease struct {
	TenantID       uint
	CreationTime   time.Time
	ExpirationTime time.Time
}

func NewLease(id uint, duration time.Duration) *Lease {
	now := time.Now()
	return &Lease{
		TenantID:       id,
		CreationTime:   now,
		ExpirationTime: now.Add(duration),
	}
}

func (l *Lease) Expired() bool {
	now := time.Now()
	return now.Sub(l.ExpirationTime) > 0
}

// MarshalBinary is an implementation of encoding.BinaryMarshaler
func (l *Lease) MarshalBinary() (data []byte, err error) {
	return json.Marshal(*l)
}

// UnmarshalBinary is an implementation of encoding.BinaryUnmarshaler
func (l *Lease) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
