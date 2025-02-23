package randx

import (
	"math/rand"
	"time"
)

func TimeDuration4Sleep(s ...time.Duration) time.Duration {
	start, end := time.Second, time.Second*3
	if len(s) > 1 {
		start = s[0]
		end = s[1]
	} else if len(s) > 0 {
		start = 0
		end = s[0]
	}

	return time.Duration(rand.Int63n(int64(end-start)) + int64(start))
}
