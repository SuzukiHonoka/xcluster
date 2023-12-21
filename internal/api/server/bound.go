package server

import "xcluster/internal/server"

type BoundInfo struct {
	ServerID  server.ID `json:"id"`
	SecretRaw string    `json:"secret"`
	Endpoint  string    `json:"endpoint,omitempty"` // configurable
}

func NewBound(id server.ID, secret, endpoint string) BoundInfo {
	return BoundInfo{
		ServerID:  id,
		SecretRaw: secret,
		Endpoint:  endpoint,
	}
}
