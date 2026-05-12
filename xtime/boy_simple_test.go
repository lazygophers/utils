package xtime

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkBeginningOfYear_Optimized(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BeginningOfYear()
	}
}

func TestBeginningOfYear_Performance(t *testing.T) {
	iterations := 1000000

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = BeginningOfYear()
	}
	duration := time.Since(start)

	fmt.Printf("\nBeginningOfYear Performance:\n")
	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Total: %v\n", duration)
	fmt.Printf("Avg: %d ns/op\n", duration.Nanoseconds()/int64(iterations))
}
