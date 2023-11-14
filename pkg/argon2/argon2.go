package argon2

// inspirations from https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

import (
	"golang.org/x/crypto/argon2"
)

type Argon2 struct {
	*Params
}

func NewArgon2(params *Params) Argon2 {
	if params == nil {
		params = ParamDefault
	}
	return Argon2{params}
}

func (r Argon2) GenerateHashFromString(s string) (*HashWrapper, error) {
	return r.GenerateHash([]byte(s))
}

func (r Argon2) GenerateHash(data []byte) (*HashWrapper, error) {
	salt, err := GenerateRandomBytes(r.SaltLength)
	if err != nil {
		return nil, err
	}
	hash := argon2.IDKey(data, salt, r.Iterations, r.Memory, r.Parallelism, r.KeyLength)
	return NewHashWrapper(salt, hash, r.Params), nil
}
