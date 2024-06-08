package randx

import (
	"math/rand"
	"time"
)

func Intn(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

func Int() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

// IntnRange
// [min, max]
func IntnRange(min, max int) int {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min+1)
}

func Int64n(n int64) int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(n)
}

func Int64() int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
}

// Int64nRange
// [min, max]
func Int64nRange(min, max int64) int64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(max-min+1)
}

func Float64() float64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
}

// Float64Range
// [min, max]
func Float64Range(min, max float64) float64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*(max-min+1)
}

func Float32() float32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Float32()
}

// Float32Range
// [min, max]
func Float32Range(min, max float32) float32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Float32()*(max-min+1)
}

func Uint32() uint32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
}

// Uint32Range
// [min, max]
func Uint32Range(min, max uint32) uint32 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()%(max-min+1)
}

func Uint64() uint64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
}

// Uint64Range
// [min, max]
func Uint64Range(min, max uint64) uint64 {
	if min > max {
		return 0
	} else if min == max {
		return min
	}

	return min + rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()%(max-min+1)
}
