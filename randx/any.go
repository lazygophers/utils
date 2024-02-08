package randx

import (
	"math/rand"
	"time"
)

func Choose[T any](s []T) T {
	if len(s) == 0 {
		return *new(T)
	}

	if len(s) == 1 {
		return s[0]
	}

	return s[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(s))]
}
