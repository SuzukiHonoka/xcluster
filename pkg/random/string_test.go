package random

import "testing"

var (
	n       = 100
	length  = 16
	strings = make([]string, n)
)

func TestRandomString(t *testing.T) {
	//strings := make([]string, n)
	for i := range strings {
		s := String(length)
		strings[i] = s
		t.Log("random string:", s)
	}
}

func TestRandomStringRepeat(t *testing.T) {
	//strings := make([]string, n)
	for _, s := range strings {
		count := -1
		for _, ss := range strings {
			if ss == s {
				count++
			}
		}
		if count > 0 {
			t.Fatal("random string repeated")
		}
	}
}
