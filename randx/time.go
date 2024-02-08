package randx

import (
	"math/rand"
	"time"
)

func TimeDuration4Sleep(s ...time.Duration) time.Duration {
	start, end := time.Second, time.Second*3
	if len(s) > 0 {
		start = s[0]
	}

	if len(s) > 1 {
		end = s[1]
	}

	return time.Duration(rand.Int63n(int64(end-start)) + int64(start))
}
