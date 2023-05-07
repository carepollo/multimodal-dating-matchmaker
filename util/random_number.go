package util

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max)
}
