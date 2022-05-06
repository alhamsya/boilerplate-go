package utils

import (
	"math/rand"
	"time"
)

func stringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString ...
func (l *thing) RandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return stringWithCharset(length, charset)
}
