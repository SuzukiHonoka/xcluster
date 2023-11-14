package argon2

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type HashWrapper struct {
	Salt   []byte
	Hash   []byte
	Params *Params
}

func NewHashWrapper(salt, hash []byte, params *Params) *HashWrapper {
	return &HashWrapper{
		Salt:   salt,
		Hash:   hash,
		Params: params,
	}
}

func NewHashWrapperFromString(s string) (*HashWrapper, error) {
	slots := strings.Split(s, "$")
	if len(slots) != 6 {
		return nil, ErrInvalidFormat
	}
	if slots[1] != "argon2id" {
		return nil, ErrNotSupported
	}
	var err error
	var version int
	if _, err = fmt.Sscanf(slots[2], "v=%d", &version); err != nil {
		return nil, err
	}
	if version != argon2.Version {
		return nil, ErrIncompatibleVersion
	}
	p := new(Params)
	if _, err = fmt.Sscanf(slots[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism); err != nil {
		return nil, err
	}
	wrapper := new(HashWrapper)
	if wrapper.Salt, err = base64.RawStdEncoding.Strict().DecodeString(slots[4]); err != nil {
		return nil, err
	}
	p.SaltLength = uint32(len(wrapper.Salt))
	if wrapper.Hash, err = base64.RawStdEncoding.Strict().DecodeString(slots[5]); err != nil {
		return nil, err
	}
	p.KeyLength = uint32(len(wrapper.Hash))
	wrapper.Params = p
	return wrapper, nil
}

func (h *HashWrapper) String() string {
	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(h.Salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(h.Hash)
	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version,
		h.Params.Memory, h.Params.Iterations, h.Params.Parallelism, b64Salt, b64Hash)
	return encodedHash
}

func (h *HashWrapper) CompareString(s string) bool {
	return h.Compare([]byte(s))
}

func (h *HashWrapper) Compare(data []byte) bool {
	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey(data,
		h.Salt, h.Params.Iterations, h.Params.Memory, h.Params.Parallelism, h.Params.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	return subtle.ConstantTimeCompare(h.Hash, otherHash) == 1
}
