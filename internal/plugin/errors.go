package plugin

import "errors"

var (
	ErrPayloadValueNotFound     = errors.New("payload value not found")
	ErrPayloadValueTypeMismatch = errors.New("payload value type mismatch")
)
