package util

import (
	"math/rand"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}
