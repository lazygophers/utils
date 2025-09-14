package fake

import (
	"runtime"
	"sync"
	"testing"
)

// 性能优先对比测试
func BenchmarkPerformanceFirstComparison(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
	
	b.Run("UltraFast", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.UltraFastName()
		}
	})
	
	b.Run("PrecomputedFast", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.PrecomputedFastName()
		}
	})
	
	b.Run("NoAlloc", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.NoAllocName()
		}
	})
	
	// 全局函数测试
	b.Run("GlobalUltraFast", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = UltraFastName()
		}
	})
	
	b.Run("GlobalPrecomputed", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = PrecomputedFastName()
		}
	})
}

// 批量生成性能对比
func BenchmarkBatchPerformanceFirst(b *testing.B) {
	b.Run("OriginalBatch1000", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(1000)
		}
	})
	
	b.Run("UltraFastBatch1000", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.BatchUltraFastNames(1000)
		}
	})
}

// 并发性能测试
func BenchmarkConcurrentPerformanceFirst(b *testing.B) {
	b.Run("OriginalConcurrent", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()
		
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.Name()
			}
		})
	})
	
	b.Run("UltraFastConcurrent", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.UltraFastName()
			}
		})
	})
	
	b.Run("PrecomputedConcurrent", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()
		
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.PrecomputedFastName()
			}
		})
	})
	
	b.Run("GlobalConcurrent", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = UltraFastName()
			}
		})
	})
}

// 内存使用详细对比
func BenchmarkMemoryDetailedComparison(b *testing.B) {
	b.Run("OriginalMemoryProfile", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := New()
		
		runtime.GC()
		runtime.ReadMemStats(&m1)
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
		
		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		
		b.Logf("Total Alloc: %d bytes", m2.TotalAlloc-m1.TotalAlloc)
		b.Logf("Heap Alloc: %d bytes", m2.HeapAlloc-m1.HeapAlloc)
		b.Logf("Mallocs: %d", m2.Mallocs-m1.Mallocs)
		b.Logf("Frees: %d", m2.Frees-m1.Frees)
		b.Logf("GC Cycles: %d", m2.NumGC-m1.NumGC)
	})
	
	b.Run("UltraFastMemoryProfile", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := NewPerformanceFirst()
		
		runtime.GC()
		runtime.ReadMemStats(&m1)
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.UltraFastName()
		}
		
		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		
		b.Logf("Total Alloc: %d bytes", m2.TotalAlloc-m1.TotalAlloc)
		b.Logf("Heap Alloc: %d bytes", m2.HeapAlloc-m1.HeapAlloc)
		b.Logf("Mallocs: %d", m2.Mallocs-m1.Mallocs)
		b.Logf("Frees: %d", m2.Frees-m1.Frees)
		b.Logf("GC Cycles: %d", m2.NumGC-m1.NumGC)
	})
	
	b.Run("PrecomputedMemoryProfile", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := NewPerformanceFirst()
		
		runtime.GC()
		runtime.ReadMemStats(&m1)
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.PrecomputedFastName()
		}
		
		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		
		b.Logf("Total Alloc: %d bytes", m2.TotalAlloc-m1.TotalAlloc)
		b.Logf("Heap Alloc: %d bytes", m2.HeapAlloc-m1.HeapAlloc)
		b.Logf("Mallocs: %d", m2.Mallocs-m1.Mallocs)
		b.Logf("Frees: %d", m2.Frees-m1.Frees)
		b.Logf("GC Cycles: %d", m2.NumGC-m1.NumGC)
	})
}

// 极限压力测试
func BenchmarkExtremePressure(b *testing.B) {
	b.Run("OriginalExtreme", func(b *testing.B) {
		faker := New()
		var wg sync.WaitGroup
		const numGoroutines = 1000
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.Name()
				}
			}()
		}
		wg.Wait()
	})
	
	b.Run("UltraFastExtreme", func(b *testing.B) {
		faker := NewPerformanceFirst()
		var wg sync.WaitGroup
		const numGoroutines = 1000
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.UltraFastName()
				}
			}()
		}
		wg.Wait()
	})
	
	b.Run("PrecomputedExtreme", func(b *testing.B) {
		faker := NewPerformanceFirst()
		var wg sync.WaitGroup
		const numGoroutines = 1000
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.PrecomputedFastName()
				}
			}()
		}
		wg.Wait()
	})
}

// CPU效率测试
func BenchmarkCPUEfficiencyExtreme(b *testing.B) {
	// 单线程CPU效率
	b.Run("OriginalCPU", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
	
	b.Run("UltraFastCPU", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			_ = faker.UltraFastName()
		}
	})
	
	b.Run("PrecomputedCPU", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			_ = faker.PrecomputedFastName()
		}
	})
}

// 随机度快速检查
func BenchmarkRandomnessQuickCheck(b *testing.B) {
	const sampleSize = 10000
	
	b.Run("UltraFastRandomness", func(b *testing.B) {
		faker := NewPerformanceFirst()
		results := make(map[string]int)
		
		b.ResetTimer()
		for i := 0; i < sampleSize; i++ {
			name := faker.UltraFastName()
			results[name]++
		}
		b.StopTimer()
		
		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)", 
			uniqueCount, sampleSize, float64(uniqueCount)/float64(sampleSize)*100)
	})
	
	b.Run("PrecomputedRandomness", func(b *testing.B) {
		faker := NewPerformanceFirst()
		results := make(map[string]int)
		
		b.ResetTimer()
		for i := 0; i < sampleSize; i++ {
			name := faker.PrecomputedFastName()
			results[name]++
		}
		b.StopTimer()
		
		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)", 
			uniqueCount, sampleSize, float64(uniqueCount)/float64(sampleSize)*100)
	})
}