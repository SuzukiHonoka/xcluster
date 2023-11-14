package argon2

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidFormat       = errors.New("invalid argon2 hash format")
	ErrNotSupported        = errors.New("hash type not supported")
	ErrIncompatibleVersion = fmt.Errorf("argon2 version incompatible, current=%d", argon2.Version)
)
