# Optimal Cache (Belady's Algorithm)

A thread-safe implementation of Belady's optimal cache replacement algorithm that provides theoretical maximum hit rates by using future knowledge of access patterns. This cache is primarily used for analysis, simulation, and benchmarking other cache algorithms.

## Features

- **Theoretical optimality** - mathematically proven best possible hit rates
- **Future access pattern** support for simulation scenarios
- **O(1)** operations with optimal eviction decisions
- **Thread-safe** implementation for concurrent simulation
- **Generic support** for any comparable key and any value type
- **Simulation framework** for cache algorithm comparison
- **Detailed statistics** for performance analysis

## How Optimal Cache Works

Belady's algorithm works by:

1. **Future Knowledge**: Requires complete future access pattern
2. **Optimal Decisions**: Always evicts the item that will be accessed farthest in the future
3. **Never-accessed Items**: Items never accessed again are evicted first
4. **Perfect Information**: Makes decisions based on complete future knowledge

This provides the theoretical upper bound on cache performance that no other algorithm can exceed.

## Use Cases

- **Algorithm benchmarking** - establishing performance baselines
- **Cache simulation studies** comparing different algorithms
- **Theoretical analysis** of workload cachability
- **Research and development** of new cache algorithms
- **Performance analysis** determining maximum possible improvements
- **Workload characterization** understanding access pattern implications

## API

```go
// Create optimal cache
cache := optimal.New[string, int](capacity)

// With known access pattern for simulation
pattern := []string{"a", "b", "c", "a", "b", "d"}
cache := optimal.NewWithPattern[string, int](capacity, pattern)

// With eviction callback
cache := optimal.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations (advances simulation time)
cache.Put("key", 42)           // Add item and advance time
value, ok := cache.Get("key")  // Retrieve and advance time
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without advancing time

// Simulation features
cache.SetAccessPattern(pattern) // Set future access pattern
currentTime := cache.CurrentTime() // Get simulation time
stats := cache.Simulate(operations) // Run full simulation

// Management
cache.Clear()                  // Reset cache and time
cache.Resize(newSize)         // Change capacity
cache.Len()                   // Current item count
cache.Cap()                   // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/optimal"
)

func main() {
    // Define access pattern for simulation
    pattern := []string{"A", "B", "C", "A", "D", "E", "A"}
    
    // Create optimal cache with known pattern
    cache := optimal.NewWithPattern[string, int](2, pattern)
    
    fmt.Println("Simulating optimal cache decisions:")
    
    // Add initial items
    cache.Put("A", 1)  // Next access at position 3
    cache.Put("B", 2)  // Next access never
    
    fmt.Printf("After adding A,B: %v\n", cache.Keys())
    
    // Add C - should evict B (never accessed again)
    cache.Put("C", 3)  // Next access never
    fmt.Printf("After adding C: %v\n", cache.Keys())
    
    // Access A (position 3 in pattern)
    cache.Get("A")
    fmt.Printf("After accessing A: %v\n", cache.Keys())
    
    // Add D - should evict C (C accessed never, A accessed at pos 6)
    cache.Put("D", 4)  // Next access never
    fmt.Printf("After adding D: %v\n", cache.Keys())
    
    // Add E - should evict D (D never accessed again)  
    cache.Put("E", 5)
    fmt.Printf("After adding E: %v\n", cache.Keys())
    
    // Final access to A
    cache.Get("A")
    fmt.Printf("Final state: %v\n", cache.Keys())
}
```

## Simulation Framework

### Operation Types

```go
// Define operations for simulation
operations := []optimal.Operation[string, int]{
    {Type: optimal.OpPut, Key: "A", Value: 1},
    {Type: optimal.OpGet, Key: "A"},
    {Type: optimal.OpPut, Key: "B", Value: 2},
    {Type: optimal.OpGet, Key: "B"},
}

// Run simulation
stats := cache.Simulate(operations)
```

### Complete Simulation Example

```go
func simulateWorkload() {
    cache := optimal.New[string, int](100)
    
    // Create realistic workload
    operations := []optimal.Operation[string, int]{}
    
    // Add initial data
    for i := 0; i < 100; i++ {
        op := optimal.Operation[string, int]{
            Type:  optimal.OpPut,
            Key:   fmt.Sprintf("item_%d", i),
            Value: i,
        }
        operations = append(operations, op)
    }
    
    // Add access pattern with some repeats
    for i := 0; i < 1000; i++ {
        op := optimal.Operation[string, int]{
            Type: optimal.OpGet,
            Key:  fmt.Sprintf("item_%d", rand.Intn(100)),
        }
        operations = append(operations, op)
    }
    
    // Run simulation
    stats := cache.Simulate(operations)
    
    fmt.Printf("Optimal Performance:\n")
    fmt.Printf("  Hits: %d\n", stats.Hits)
    fmt.Printf("  Misses: %d\n", stats.Misses)
    fmt.Printf("  Hit Rate: %.2f%%\n", stats.HitRate*100)
    fmt.Printf("  Evictions: %d\n", stats.Evictions)
}
```

## Performance

- **Get**: O(1) - Hash lookup + future access time lookup
- **Put**: O(1) - Hash insert + optimal eviction decision
- **Remove**: O(1) - Hash delete + access pattern cleanup
- **Eviction**: O(n) - Scan for farthest future access (amortized O(1))
- **Memory**: O(n + p) where n=items, p=pattern length

## Benchmarks

```
BenchmarkOptimalPut-8      3000000    400 ns/op    88 B/op    2 allocs/op
BenchmarkOptimalGet-8      5000000    300 ns/op    24 B/op    1 allocs/op
BenchmarkSimulate-8           1000   1200000 ns/op  45000 B/op  1500 allocs/op
```

## Statistics

Comprehensive analysis of optimal performance:

```go
stats := cache.Simulate(operations)

fmt.Printf("Optimal Cache Analysis:\n")
fmt.Printf("  Cache Hits: %d\n", stats.Hits)
fmt.Printf("  Cache Misses: %d\n", stats.Misses)
fmt.Printf("  Total Operations: %d\n", stats.Hits + stats.Misses)
fmt.Printf("  Hit Rate: %.2f%%\n", stats.HitRate * 100)
fmt.Printf("  Evictions: %d\n", stats.Evictions)

// This represents the theoretical maximum hit rate
fmt.Printf("  Theoretical Maximum Hit Rate: %.2f%%\n", stats.HitRate * 100)
```

## Algorithm Comparison

Use optimal cache to benchmark other algorithms:

```go
func compareAlgorithms(operations []Operation[string, int]) {
    // Run optimal simulation
    optimalCache := optimal.New[string, int](100)
    optimalStats := optimalCache.Simulate(operations)
    
    // Run LRU simulation (conceptual)
    // lruStats := simulateLRU(operations, 100)
    
    fmt.Printf("Algorithm Comparison:\n")
    fmt.Printf("  Optimal Hit Rate: %.2f%%\n", optimalStats.HitRate*100)
    // fmt.Printf("  LRU Hit Rate: %.2f%%\n", lruStats.HitRate*100)
    // fmt.Printf("  Improvement Potential: %.2f%%\n", 
    //     (optimalStats.HitRate - lruStats.HitRate)*100)
}
```

## Limitations

### Practical Limitations
- **Future knowledge required**: Cannot be used in real applications
- **Access pattern dependency**: Performance tied to specific patterns
- **Memory overhead**: Must store entire access pattern
- **Simulation only**: Not suitable for production caching

### Theoretical Insights
- **Upper bound**: No algorithm can exceed optimal performance
- **Workload analysis**: Shows maximum possible cache effectiveness
- **Algorithm gaps**: Reveals how much improvement is theoretically possible

## When to Use Optimal Cache

### ✅ Perfect for:
- **Algorithm research** developing new cache replacement strategies
- **Performance benchmarking** establishing theoretical limits
- **Simulation studies** comparing algorithm effectiveness
- **Workload analysis** understanding access pattern characteristics
- **Academic research** in caching and memory management

### ✅ Use for analysis when:
- **Evaluating cache algorithms** against theoretical optimum
- **Sizing cache capacity** based on workload characteristics
- **Research and development** of caching systems
- **Understanding workload** cachability and temporal patterns

### ❌ Cannot be used for:
- **Production applications** (requires future knowledge)
- **Real-time systems** (simulation only)
- **Online caching** (no access pattern prediction)
- **General-purpose caching** (use practical algorithms)

## Research Applications

### Cache Algorithm Development
```go
// Test new algorithm against optimal
func evaluateNewAlgorithm(workload []Operation[string, int]) float64 {
    optimalStats := optimal.Simulate(workload)
    newAlgorithmStats := newAlgorithm.Simulate(workload)
    
    // Calculate efficiency: how close to optimal
    efficiency := newAlgorithmStats.HitRate / optimalStats.HitRate
    return efficiency // Value between 0.0 and 1.0
}
```

### Workload Characterization
```go
// Analyze workload cachability
func analyzeWorkload(operations []Operation[string, int], capacities []int) {
    for _, capacity := range capacities {
        cache := optimal.New[string, int](capacity)
        stats := cache.Simulate(operations)
        
        fmt.Printf("Capacity %d: %.2f%% hit rate\n", 
            capacity, stats.HitRate*100)
    }
}
```

Belady's optimal cache provides the theoretical foundation for understanding cache performance limits and serves as the gold standard for evaluating and developing cache replacement algorithms.