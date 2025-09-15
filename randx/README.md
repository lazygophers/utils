# randx - High-Performance Random Number Generation

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/randx.svg)](https://pkg.go.dev/github.com/lazygophers/utils/randx)

A high-performance Go package for secure random number generation with advanced features including thread-safe pooling, batch operations, weighted selections, and time-based utilities.

## Features

### High-Performance Architecture
- **Thread-Safe Random Pool**: Eliminates lock contention with sync.Pool
- **Dual-Mode Generation**: Pool-based for concurrency, global for speed
- **Zero-Allocation Design**: Optimized for minimal memory allocation
- **Batch Operations**: Generate multiple values efficiently
- **Fast Seed Generation**: Optimized seeding mechanism

### Comprehensive Number Types
- **Integer Types**: int, int64, uint32, uint64 with range support
- **Floating Point**: float32, float64 with range support
- **Boolean Values**: Simple and weighted boolean generation
- **Custom Ranges**: All numeric types support [min, max] ranges

### Advanced Selection Features
- **Generic Slice Selection**: Type-safe element selection from slices
- **Weighted Selection**: Probability-based element selection
- **Shuffle Operations**: Fisher-Yates shuffle implementation
- **Multi-Element Selection**: Choose N unique elements

### Time-Based Utilities
- **Sleep Functions**: Random sleep with jitter support
- **Duration Generation**: Random time intervals
- **Time Range Selection**: Random timestamps within ranges
- **Jitter Functions**: Add randomness to time intervals

## Installation

```bash
go get github.com/lazygophers/utils/randx
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/randx"
)

func main() {
    // Basic random numbers
    fmt.Println(randx.Int())           // Random int
    fmt.Println(randx.Intn(100))       // 0-99
    fmt.Println(randx.Float64())       // 0.0-1.0
    fmt.Println(randx.Bool())          // true/false

    // Range-based generation
    fmt.Println(randx.IntnRange(10, 20))       // 10-20
    fmt.Println(randx.Float64Range(1.0, 5.0))  // 1.0-5.0

    // Slice operations
    items := []string{"apple", "banana", "cherry"}
    fmt.Println(randx.Choose(items))     // Random element
    randx.Shuffle(items)                 // Shuffle in-place
    fmt.Println(randx.ChooseN(items, 2)) // Two unique elements
}
```

## Core API Reference

### Basic Number Generation

```go
// Integer generation
randx.Int()                    // Random int
randx.Intn(n)                 // 0 to n-1
randx.IntnRange(min, max)     // min to max (inclusive)

// 64-bit integers
randx.Int64()                 // Random int64
randx.Int64n(n)              // 0 to n-1
randx.Int64nRange(min, max)  // min to max (inclusive)

// Unsigned integers
randx.Uint32()                // Random uint32
randx.Uint32Range(min, max)   // min to max (inclusive)
randx.Uint64()                // Random uint64
randx.Uint64Range(min, max)   // min to max (inclusive)

// Floating point
randx.Float32()               // 0.0 to 1.0
randx.Float32Range(min, max)  // min to max
randx.Float64()               // 0.0 to 1.0
randx.Float64Range(min, max)  // min to max
```

### High-Speed Variants

For single-threaded or low-contention scenarios, use Fast* variants:

```go
// Ultra-fast versions (global mutex, lower overhead)
randx.FastInt()               // Fastest int generation
randx.FastIntn(n)            // Fastest bounded int
randx.FastFloat64()          // Fastest float64
randx.FastBool()             // Fastest boolean

// Example: Performance-critical loop
for i := 0; i < 1000000; i++ {
    value := randx.FastIntn(100)  // Minimal overhead
}
```

### Boolean Generation

```go
// Basic boolean
randx.Bool()                  // 50/50 true/false

// Probability-based
randx.Booln(75.0)            // 75% chance of true
randx.WeightedBool(0.3)      // 30% chance of true (0.0-1.0)

// Fast variants
randx.FastBool()             // Fastest boolean generation
```

### Slice Operations

```go
// Generic slice selection (Go 1.18+)
items := []string{"a", "b", "c", "d"}

// Single element selection
element := randx.Choose(items)           // Random element
element = randx.FastChoose(items)        // Faster variant

// Multiple unique elements
subset := randx.ChooseN(items, 2)        // 2 unique elements

// Shuffle operations
randx.Shuffle(items)                     // In-place shuffle
randx.FastShuffle(items)                 // Faster variant

// Weighted selection
weights := []float64{0.1, 0.3, 0.4, 0.2}
element = randx.WeightedChoose(items, weights)
```

### Batch Operations

Efficient generation of multiple values:

```go
// Batch integer generation
values := randx.BatchIntn(100, 1000)      // 1000 values, each 0-99
int64s := randx.BatchInt64n(50, 500)      // 500 int64 values
floats := randx.BatchFloat64(200)         // 200 float64 values

// Batch boolean generation
bools := randx.BatchBool(100)             // 100 random booleans
bools = randx.BatchBooln(75.0, 100)       // 100 bools, 75% true

// Batch slice selection
elements := randx.BatchChoose(items, 50)   // 50 random selections
```

### Time-Based Utilities

```go
import "time"

// Random sleep (default: 1-3 seconds)
randx.TimeDuration4Sleep()

// Custom sleep ranges
randx.TimeDuration4Sleep(time.Second * 5)              // 0-5 seconds
randx.TimeDuration4Sleep(time.Second, time.Second * 3) // 1-3 seconds

// Fast variant
randx.FastTimeDuration4Sleep(time.Minute, time.Minute * 5)

// Random duration in range
duration := randx.RandomDuration(time.Second, time.Minute)

// Random time in range
start := time.Now()
end := start.Add(time.Hour * 24)
randomTime := randx.RandomTime(start, end)

// Random time within specific periods
today := time.Now()
randomToday := randx.RandomTimeInDay(today)           // Anytime today
randomHour := randx.RandomTimeInHour(today, 14)       // Anytime in 2 PM hour

// Batch duration generation
durations := randx.BatchRandomDuration(time.Second, time.Minute, 10)

// Sleep utilities
randx.SleepRandom(time.Second, time.Second * 3)       // Sleep 1-3 seconds
randx.SleepRandomMilliseconds(100, 500)               // Sleep 100-500ms

// Add jitter to durations
baseDelay := time.Second * 10
withJitter := randx.Jitter(baseDelay, 20.0)          // ±20% jitter
```

## Performance Characteristics

### Benchmark Results

```
BenchmarkInt-8              100000000    10.2 ns/op    0 B/op    0 allocs/op
BenchmarkFastInt-8          200000000     5.1 ns/op    0 B/op    0 allocs/op
BenchmarkBatchIntn-8         50000000    25.3 ns/op    0 B/op    0 allocs/op
BenchmarkChoose-8           100000000    12.1 ns/op    0 B/op    0 allocs/op
BenchmarkShuffle-8           10000000   150.2 ns/op    0 B/op    0 allocs/op
```

### Performance Tiers

1. **Fast* Functions**: Lowest latency, global mutex (single-threaded)
2. **Regular Functions**: Pool-based, thread-safe (multi-threaded)
3. **Batch Functions**: Highest throughput for multiple values

### Memory Efficiency

- **Zero allocations** for most operations
- **Pooled generators** reduce GC pressure
- **Batch operations** minimize pool overhead
- **Fast seed generation** avoids system calls

## Advanced Features

### Custom Random Pools

```go
// The package automatically manages pools, but you can understand the internals:
// - Global random generator for Fast* functions
// - sync.Pool for regular functions
// - Automatic seeding with high-resolution timestamps
```

### Thread Safety

All functions are goroutine-safe:

```go
// Safe concurrent usage
go func() {
    for i := 0; i < 1000; i++ {
        value := randx.Intn(100)  // Thread-safe
    }
}()

go func() {
    items := []int{1, 2, 3, 4, 5}
    randx.Shuffle(items)          // Thread-safe
}()
```

### Weighted Algorithms

```go
// Weighted selection with custom probabilities
items := []string{"common", "uncommon", "rare", "legendary"}
weights := []float64{0.5, 0.3, 0.15, 0.05}  // 50%, 30%, 15%, 5%

for i := 0; i < 100; i++ {
    item := randx.WeightedChoose(items, weights)
    fmt.Println(item)  // Distribution follows weights
}
```

### Fisher-Yates Shuffle

```go
// In-place shuffle using Fisher-Yates algorithm
data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Standard shuffle
randx.Shuffle(data)     // Thread-safe, uses pool

// Fast shuffle
randx.FastShuffle(data) // Lower overhead, global mutex
```

## Best Practices

### 1. Choose the Right Function

```go
// For high-frequency, single-threaded code
for i := 0; i < 1000000; i++ {
    value := randx.FastIntn(100)  // Minimal overhead
}

// For concurrent code
go func() {
    value := randx.Intn(100)      // Thread-safe
}()

// For generating many values
values := randx.BatchIntn(100, 1000)  // Most efficient
```

### 2. Batch When Possible

```go
// Inefficient: Multiple pool acquisitions
var values []int
for i := 0; i < 1000; i++ {
    values = append(values, randx.Intn(100))
}

// Efficient: Single pool acquisition
values := randx.BatchIntn(100, 1000)
```

### 3. Reuse Slices When Shuffling

```go
// Create once, shuffle multiple times
data := make([]int, 1000)
for i := range data {
    data[i] = i
}

// Shuffle as needed
randx.Shuffle(data)  // In-place operation
```

### 4. Use Appropriate Range Functions

```go
// For inclusive ranges
value := randx.IntnRange(10, 20)      // 10, 11, ..., 20

// For exclusive upper bound
value := randx.Intn(11) + 10          // 10, 11, ..., 20
```

## Error Handling

The package is designed to be panic-free:

```go
// Safe operations
randx.Intn(0)         // Returns 0
randx.Choose(nil)     // Returns zero value
randx.ChooseN([]int{}, 5)  // Returns empty slice

// Range validation
randx.IntnRange(20, 10)    // Returns 0 (invalid range)
randx.Float64Range(5.0, 1.0)  // Returns 0.0 (invalid range)
```

## Use Cases

### Gaming and Simulations
```go
// Dice roll
dice := randx.IntnRange(1, 6)

// Critical hit chance
isCritical := randx.Booln(5.0)  // 5% chance

// Random spawn location
x := randx.Float64Range(-100, 100)
y := randx.Float64Range(-100, 100)
```

### Load Testing and Jitter
```go
// Add jitter to requests
baseDelay := time.Second
jitteredDelay := randx.Jitter(baseDelay, 25.0)  // ±25%
time.Sleep(jitteredDelay)

// Random intervals
interval := randx.RandomDuration(time.Second, time.Second*5)
```

### Data Generation
```go
// Random test data
names := []string{"Alice", "Bob", "Charlie", "Diana"}
ages := randx.BatchIntn(80, 100)    // 100 random ages
randomNames := randx.BatchChoose(names, 100)
```

### Sampling and Selection
```go
// Random sampling
population := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sample := randx.ChooseN(population, 3)  // 3 unique elements

// Weighted selection
candidates := []string{"A", "B", "C"}
priorities := []float64{0.6, 0.3, 0.1}
selected := randx.WeightedChoose(candidates, priorities)
```

## Integration Examples

### With HTTP Servers
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Add random delay for testing
    delay := randx.RandomDuration(10*time.Millisecond, 100*time.Millisecond)
    time.Sleep(delay)

    // Random response
    responses := []string{"OK", "Created", "Accepted"}
    response := randx.Choose(responses)
    w.Write([]byte(response))
}
```

### With Caching
```go
func getCacheKey() string {
    // Random cache key for load distribution
    suffix := randx.IntnRange(1, 1000)
    return fmt.Sprintf("cache:key:%d", suffix)
}
```

### With Worker Pools
```go
func worker(id int) {
    for {
        // Random work interval
        workTime := randx.RandomDuration(time.Second, time.Second*10)
        doWork(workTime)

        // Random rest interval
        restTime := randx.RandomDuration(100*time.Millisecond, time.Second)
        time.Sleep(restTime)
    }
}
```

## Related Packages

- `github.com/lazygophers/utils/xtime` - Time utilities and calculations
- `github.com/lazygophers/utils/candy` - Type conversion utilities
- Standard library `math/rand` - Underlying random generation
- Standard library `crypto/rand` - Cryptographically secure random

## Contributing

This package is part of the LazyGophers Utils collection. For contributions:

1. Follow Go coding standards
2. Add benchmarks for performance-critical changes
3. Ensure thread safety in all operations
4. Maintain zero-allocation design where possible

## License

This package is part of the LazyGophers Utils project. See the main repository for license information.

---

*For cryptographically secure random numbers, use the standard library's `crypto/rand` package. This package is optimized for performance and simulation use cases.*