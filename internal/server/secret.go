package server

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const nonceLength = 8

type Secret []byte

func NewRandomSecret() (string, Secret, error) {
	nonce := make([]byte, nonceLength)
	if _, err := rand.Read(nonce); err != nil {
		return "", nil, err
	}
	h := sha256.New()
	h.Write(nonce)
	hash := h.Sum(nil)
	b64Nonce := base64.RawStdEncoding.EncodeToString(nonce)
	return b64Nonce, hash, nil
}

func (s Secret) Compare(other string) (bool, error) {
	b, err := base64.RawStdEncoding.DecodeString(other)
	if err != nil {
		return false, err
	}
	h := sha256.New()
	h.Write(b)
	hash := h.Sum(nil)
	return bytes.Equal(s, hash), nil
}
