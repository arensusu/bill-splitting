package helper

import (
	"math/rand"
	"time"
)

var randGen *rand.Rand

func init() {
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[randGen.Intn(len(letters))]
	}
	return string(b)
}

func RandomInt64(min, max int64) int64 {
	return min + randGen.Int63n(max-min+1)
}

func RandomDate() time.Time {
	return time.Now().AddDate(0, 0, randGen.Intn(365))
}
