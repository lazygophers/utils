package fake

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

// 极限性能对比测试
func BenchmarkExtremeComparison(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("PreviousBest_PrecomputedFast", func(b *testing.B) {
		faker := NewPerformanceFirst()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.PrecomputedFastName()
		}
	})

	b.Run("Extreme", func(b *testing.B) {
		faker := NewExtremePerformance()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ExtremeName()
		}
	})

	b.Run("ZeroAllocExtreme", func(b *testing.B) {
		faker := NewExtremePerformance()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ZeroAllocExtremeName()
		}
	})

	b.Run("UltraCompact", func(b *testing.B) {
		faker := NewUltraCompact()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.CompactName()
		}
	})

	b.Run("Inline", func(b *testing.B) {
		faker := NewInline()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.InlineName()
		}
	})

	b.Run("Assembly", func(b *testing.B) {
		faker := NewAssemblyOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.AssemblyName()
		}
	})

	b.Run("MemoryMapped", func(b *testing.B) {
		faker := NewMemoryMapped()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.MemoryMappedName()
		}
	})

	// 全局函数测试
	b.Run("GlobalExtreme", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = ExtremeName()
		}
	})

	b.Run("GlobalCompact", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = CompactName()
		}
	})

	b.Run("GlobalInline", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = InlineName()
		}
	})

	b.Run("GlobalAssembly", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = AssemblyName()
		}
	})

	b.Run("GlobalMemoryMapped", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = MemoryMappedName()
		}
	})
}

// 批量生成极限性能测试
func BenchmarkExtremeBatch(b *testing.B) {
	b.Run("Original_Batch1000", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(1000)
		}
	})

	b.Run("Extreme_Batch1000", func(b *testing.B) {
		faker := NewExtremePerformance()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.BatchExtreme(1000)
		}
	})
}

// 并发极限性能测试
func BenchmarkExtremeConcurrent(b *testing.B) {
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

	b.Run("Extreme", func(b *testing.B) {
		faker := NewExtremePerformance()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.ExtremeName()
			}
		})
	})

	b.Run("UltraCompact", func(b *testing.B) {
		faker := NewUltraCompact()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.CompactName()
			}
		})
	})

	b.Run("Assembly", func(b *testing.B) {
		faker := NewAssemblyOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.AssemblyName()
			}
		})
	})

	b.Run("MemoryMapped", func(b *testing.B) {
		faker := NewMemoryMapped()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.MemoryMappedName()
			}
		})
	})

	b.Run("GlobalExtreme", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = ExtremeName()
			}
		})
	})
}

// 内存使用极限测试
func BenchmarkExtremeMemory(b *testing.B) {
	b.Run("OriginalMemory", func(b *testing.B) {
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

		b.Logf("Memory: %d bytes, GC: %d", m2.TotalAlloc-m1.TotalAlloc, m2.NumGC-m1.NumGC)
	})

	b.Run("ExtremeMemory", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := NewExtremePerformance()

		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ExtremeName()
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)

		b.Logf("Memory: %d bytes, GC: %d", m2.TotalAlloc-m1.TotalAlloc, m2.NumGC-m1.NumGC)
	})

	b.Run("CompactMemory", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := NewUltraCompact()

		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.CompactName()
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)

		b.Logf("Memory: %d bytes, GC: %d", m2.TotalAlloc-m1.TotalAlloc, m2.NumGC-m1.NumGC)
	})

	b.Run("MemoryMappedMemory", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		faker := NewMemoryMapped()

		runtime.GC()
		runtime.ReadMemStats(&m1)

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.MemoryMappedName()
		}

		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)

		b.Logf("Memory: %d bytes, GC: %d", m2.TotalAlloc-m1.TotalAlloc, m2.NumGC-m1.NumGC)
	})
}

// CPU效率极限测试
func BenchmarkExtremeCPU(b *testing.B) {
	// 单线程CPU极限
	b.Run("ExtremeVsOriginalCPU", func(b *testing.B) {
		originalFaker := New()
		extremeFaker := NewExtremePerformance()

		b.Run("Original", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = originalFaker.Name()
			}
		})

		b.Run("Extreme", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = extremeFaker.ExtremeName()
			}
		})

		b.Run("Assembly", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = AssemblyName()
			}
		})

		b.Run("MemoryMapped", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MemoryMappedName()
			}
		})
	})
}

// 压力测试 - 极限负载
func BenchmarkExtremeStress(b *testing.B) {
	const numGoroutines = 10000

	b.Run("OriginalStress", func(b *testing.B) {
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

	b.Run("ExtremeStress", func(b *testing.B) {
		faker := NewExtremePerformance()
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = faker.ExtremeName()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("GlobalAssemblyStress", func(b *testing.B) {
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < b.N/numGoroutines; j++ {
					_ = AssemblyName()
				}
			}()
		}
		wg.Wait()
	})
}

// 随机度检查 - 极限版本
func BenchmarkExtremeRandomness(b *testing.B) {
	const sampleSize = 100000

	b.Run("ExtremeRandomness", func(b *testing.B) {
		faker := NewExtremePerformance()
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < sampleSize; i++ {
			name := faker.ExtremeName()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Extreme - Unique names: %d out of %d (%.2f%%)",
			uniqueCount, sampleSize, float64(uniqueCount)/float64(sampleSize)*100)
	})

	b.Run("CompactRandomness", func(b *testing.B) {
		faker := NewUltraCompact()
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < sampleSize; i++ {
			name := faker.CompactName()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Compact - Unique names: %d out of %d (%.2f%%)",
			uniqueCount, sampleSize, float64(uniqueCount)/float64(sampleSize)*100)
	})

	b.Run("AssemblyRandomness", func(b *testing.B) {
		results := make(map[string]int)

		b.ResetTimer()
		for i := 0; i < sampleSize; i++ {
			name := AssemblyName()
			results[name]++
		}
		b.StopTimer()

		uniqueCount := len(results)
		b.Logf("Assembly - Unique names: %d out of %d (%.2f%%)",
			uniqueCount, sampleSize, float64(uniqueCount)/float64(sampleSize)*100)
	})
}

// 纯原子操作性能测试
func BenchmarkAtomicOperations(b *testing.B) {
	b.Run("AtomicAddUint64", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&counter, 1)
		}
	})

	b.Run("AtomicAddUint32", func(b *testing.B) {
		var counter uint32
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			atomic.AddUint32(&counter, 1)
		}
	})

	b.Run("AtomicLoadStore", func(b *testing.B) {
		var value uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			atomic.StoreUint64(&value, atomic.LoadUint64(&value)+1)
		}
	})
}
