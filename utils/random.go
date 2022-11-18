package utils

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	size := len(chars)
	for i := 0; i < n; i++ {
		s := chars[rand.Intn(size)]
		sb.WriteByte(s)
	}
	return sb.String()
}
