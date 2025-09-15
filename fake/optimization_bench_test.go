package fake

import (
	"runtime"
	"sync"
	"testing"
)

// 基准测试：原版 vs 优化版
func BenchmarkNameComparison(b *testing.B) {
	// 原版
	b.Run("Original", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	// 优化版
	b.Run("Optimized", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.FastName()
		}
	})

	// 池化版本
	b.Run("Pooled", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.PooledName()
		}
	})

	// Unsafe版本
	b.Run("Unsafe", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.UnsafeName()
		}
	})

	// 超级优化版本
	b.Run("SuperOptimized", func(b *testing.B) {
		faker := NewSuperOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.SuperFastName()
		}
	})

	// 零分配版本
	b.Run("ZeroAlloc", func(b *testing.B) {
		faker := NewSuperOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ZeroAllocName()
		}
	})
}

// 并发性能对比
func BenchmarkConcurrentComparison(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.Name()
			}
		})
	})

	b.Run("Optimized", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.FastName()
			}
		})
	})

	b.Run("SuperOptimized", func(b *testing.B) {
		faker := NewSuperOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.SuperFastName()
			}
		})
	})
}

// 批量生成性能对比
func BenchmarkBatchComparison(b *testing.B) {
	counts := []int{10, 100, 1000}

	for _, count := range counts {
		b.Run("Original_"+string(rune(count+'0')), func(b *testing.B) {
			faker := New()
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = faker.BatchNames(count)
			}
		})

		b.Run("Optimized_"+string(rune(count+'0')), func(b *testing.B) {
			faker := NewOptimized()
			b.ResetTimer()
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				_ = faker.BatchFastNames(count)
			}
		})
	}
}

// 内存使用对比
func BenchmarkMemoryComparison(b *testing.B) {
	b.Run("OriginalMemory", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)

		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		b.Logf("Memory: %d bytes, Allocs: %d", m2.TotalAlloc-m1.TotalAlloc, m2.Mallocs-m1.Mallocs)
	})

	b.Run("OptimizedMemory", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)

		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.FastName()
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		b.Logf("Memory: %d bytes, Allocs: %d", m2.TotalAlloc-m1.TotalAlloc, m2.Mallocs-m1.Mallocs)
	})
}

// 随机度测试
func BenchmarkRandomnessComparison(b *testing.B) {
	const iterations = 100000

	b.Run("OriginalRandomness", func(b *testing.B) {
		faker := New()
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < iterations; i++ {
			name := faker.Name()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)",
			uniqueCount, iterations, float64(uniqueCount)/float64(iterations)*100)
	})

	b.Run("OptimizedRandomness", func(b *testing.B) {
		faker := NewOptimized()
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < iterations; i++ {
			name := faker.FastName()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)",
			uniqueCount, iterations, float64(uniqueCount)/float64(iterations)*100)
	})

	b.Run("SuperOptimizedRandomness", func(b *testing.B) {
		faker := NewSuperOptimized()
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < iterations; i++ {
			name := faker.SuperFastName()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)",
			uniqueCount, iterations, float64(uniqueCount)/float64(iterations)*100)
	})
}

// 锁竞争对比测试
func BenchmarkLockContentionComparison(b *testing.B) {
	const numGoroutines = 100

	b.Run("OriginalContention", func(b *testing.B) {
		faker := New()
		var wg sync.WaitGroup

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

	b.Run("OptimizedContention", func(b *testing.B) {
		faker := NewOptimized()
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.FastName()
				}
			}()
		}
		wg.Wait()
	})
}

// CPU效率测试
func BenchmarkCPUEfficiencyComparison(b *testing.B) {
	b.Run("OriginalCPU", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("OptimizedCPU", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.FastName()
		}
	})

	b.Run("SuperOptimizedCPU", func(b *testing.B) {
		faker := NewSuperOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.SuperFastName()
		}
	})
}

// 多语言性能对比
func BenchmarkMultiLanguageOptimized(b *testing.B) {
	b.Run("English", func(b *testing.B) {
		faker := NewOptimized(WithLanguage(LanguageEnglish))
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.FastName()
		}
	})

	b.Run("Chinese", func(b *testing.B) {
		faker := NewOptimized(WithLanguage(LanguageChineseSimplified))
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.FastName()
		}
	})
}

// 大规模压力测试
func BenchmarkStressTest(b *testing.B) {
	b.Run("OriginalStress", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			names := faker.BatchNames(1000)
			if len(names) != 1000 {
				b.Fatal("Batch generation failed")
			}
		}
	})

	b.Run("OptimizedStress", func(b *testing.B) {
		faker := NewOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			names := faker.BatchFastNames(1000)
			if len(names) != 1000 {
				b.Fatal("Batch generation failed")
			}
		}
	})
}
