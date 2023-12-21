package server

import (
	"crypto/sha256"
	"encoding/hex"
	"xcluster/pkg/random"
)

var nonceLength = 16 // adjustable

type Secret string // random string

func GenerateSha256String(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func NewRandomSecret() (string, Secret) {
	s := random.ComplexString(nonceLength)
	hash := GenerateSha256String(s)
	return s, Secret(hash)
}

func (s Secret) Compare(other string) bool {
	hash := GenerateSha256String(other)
	return string(s) == hash
}
