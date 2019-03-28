package helper

import (
	"math/rand"
	"sync"
	"time"
)

const source = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var size int

// generates a string of size, size filled with random ascii characters
func GenerateString(targetSize int) string {
	var once sync.Once
	once.Do(func() {
		size = len(source)
	})

	r := make([]rune, targetSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < targetSize; i++ {
		r[i] = rune(source[rand.Intn(size)])
	}
	return string(r)
}
