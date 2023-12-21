package random

import (
	"math/rand"
)

var (
	basicTable   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	complexTable = append(basicTable, []rune("~!@#$%^&*()_+`-=")...)
)

func RandString(table []rune, n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = table[rand.Intn(len(table))]
	}
	return string(s)
}

func BasicString(n int) string {
	return RandString(basicTable, n)
}

func ComplexString(n int) string {
	return RandString(complexTable, n)
}

func String(n int) string {
	return BasicString(n)
}
