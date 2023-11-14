package argon2

import (
	"math/rand"
	"testing"
)

var (
	n      = 10
	length = 32
	hashes = make(map[string][]byte, n)
)

func generateHash(n, length int, hook func(p []byte, h string)) {
	for i := 0; i < n; i++ {
		nonce, _ := GenerateRandomBytes(uint32(rand.Intn(length-1) + 1))
		argon2 := NewArgon2(nil)
		w, err := argon2.GenerateHash(nonce)
		if err != nil {
			panic(err)
		}
		hook(nonce, w.String())
	}
}

func TestArgon2_Generate(t *testing.T) {
	test := func(p []byte, h string) {
		hashes[h] = p
		t.Logf("test hash=%s", h)
	}
	generateHash(n, length, test)
}

func TestArgon2_Compare(t *testing.T) {
	for h, p := range hashes {
		t.Logf("test hash=%s", h)
		w, err := NewHashWrapperFromString(h)
		if err != nil {
			t.Error(err)
		}
		if !w.Compare(p) {
			t.Fatal("compare hash failed")
		}
	}
}
