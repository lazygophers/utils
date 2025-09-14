package fake

import (
	"runtime"
	"sync"
	"testing"
)

// 基准测试 - 基本功能
func BenchmarkNamePerformance(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

func BenchmarkEmail(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.Email()
	}
}

func BenchmarkPhoneNumber(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.PhoneNumber()
	}
}

func BenchmarkUserAgent(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.UserAgent()
	}
}

func BenchmarkAddress(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.FullAddress()
	}
}

func BenchmarkDeviceInfo(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.DeviceInfo()
	}
}

// 基准测试 - 全局函数 vs 实例方法
func BenchmarkGlobalName(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = Name()
	}
}

func BenchmarkInstanceName(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

// 基准测试 - 批量生成对比
func BenchmarkBatchNamesComparison(b *testing.B) {
	faker := New()
	
	b.Run("Batch10", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(10)
		}
	})
	
	b.Run("Batch100", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(100)
		}
	})
	
	b.Run("Batch1000", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.BatchNames(1000)
		}
	})
}

// 基准测试 - 并发性能
func BenchmarkConcurrentName(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.Name()
		}
	})
}

func BenchmarkConcurrentEmail(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.Email()
		}
	})
}

func BenchmarkConcurrentUserAgent(b *testing.B) {
	faker := New()
	b.ResetTimer()
	b.ReportAllocs()
	
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = faker.UserAgent()
		}
	})
}

// 基准测试 - 内存分配
func BenchmarkMemoryAllocation(b *testing.B) {
	faker := New()
	
	b.Run("SingleInstance", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
			_ = faker.Email()
			_ = faker.PhoneNumber()
		}
	})
	
	b.Run("MultipleInstances", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			f := New()
			_ = f.Name()
			_ = f.Email()
			_ = f.PhoneNumber()
		}
	})
}

// 基准测试 - 缓存效果
func BenchmarkWithCache(b *testing.B) {
	faker := New()
	
	// 预热缓存
	for i := 0; i < 100; i++ {
		_ = faker.Name()
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

func BenchmarkWithoutCache(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		faker := New()
		_ = faker.Name()
	}
}

// 基准测试 - 不同语言
func BenchmarkMultiLanguage(b *testing.B) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}
	
	for _, lang := range languages {
		b.Run(string(lang), func(b *testing.B) {
			faker := New(WithLanguage(lang))
			b.ResetTimer()
			b.ReportAllocs()
			
			for i := 0; i < b.N; i++ {
				_ = faker.Name()
			}
		})
	}
}

// 压力测试 - 内存使用监控
func BenchmarkMemoryStress(b *testing.B) {
	faker := New()
	
	b.Run("LargeDataGeneration", func(b *testing.B) {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)
		
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			// 生成大量数据
			_ = faker.BatchNames(100)
			_ = faker.BatchEmails(100)
			_ = faker.BatchUserAgents(100)
		}
		
		b.StopTimer()
		runtime.GC()
		runtime.ReadMemStats(&m2)
		
		b.Logf("Memory allocated: %d bytes", m2.TotalAlloc-m1.TotalAlloc)
		b.Logf("Memory used: %d bytes", m2.Alloc-m1.Alloc)
		b.Logf("GC cycles: %d", m2.NumGC-m1.NumGC)
	})
}

// 基准测试 - CPU使用优化
func BenchmarkCPUEfficiency(b *testing.B) {
	faker := New()
	
	b.Run("StandardGeneration", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
			_ = faker.Email()
			_ = faker.PhoneNumber()
			_ = faker.UserAgent()
		}
	})
	
	b.Run("BatchGeneration", func(b *testing.B) {
		b.ReportAllocs()
		batchSize := b.N / 4
		if batchSize <= 0 {
			batchSize = 1
		}
		
		for i := 0; i < 4; i++ {
			_ = faker.BatchNames(batchSize)
			_ = faker.BatchEmails(batchSize)
			_ = faker.BatchUserAgents(batchSize)
		}
	})
}

// 随机度测试
func BenchmarkRandomness(b *testing.B) {
	faker := New()
	
	b.Run("NameRandomness", func(b *testing.B) {
		results := make(map[string]int)
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			name := faker.Name()
			results[name]++
		}
		
		b.StopTimer()
		uniqueCount := len(results)
		b.Logf("Unique names: %d out of %d (%.2f%%)", 
			uniqueCount, b.N, float64(uniqueCount)/float64(b.N)*100)
	})
	
	b.Run("EmailRandomness", func(b *testing.B) {
		results := make(map[string]int)
		b.ResetTimer()
		
		for i := 0; i < b.N; i++ {
			email := faker.Email()
			results[email]++
		}
		
		b.StopTimer()
		uniqueCount := len(results)
		b.Logf("Unique emails: %d out of %d (%.2f%%)", 
			uniqueCount, b.N, float64(uniqueCount)/float64(b.N)*100)
	})
}

// 锁竞争测试
func BenchmarkLockContention(b *testing.B) {
	faker := New()
	
	b.Run("NoContention", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
	
	b.Run("HighContention", func(b *testing.B) {
		const numGoroutines = 100
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
}

// 基准测试 - 不同配置下的性能
func BenchmarkConfigurationImpact(b *testing.B) {
	b.Run("DefaultConfig", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
	
	b.Run("WithSeed", func(b *testing.B) {
		faker := New(WithSeed(12345))
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
	
	b.Run("WithMultipleOptions", func(b *testing.B) {
		faker := New(
			WithLanguage(LanguageChineseSimplified),
			WithCountry(CountryChina),
			WithGender(GenderFemale),
			WithSeed(12345),
		)
		b.ResetTimer()
		b.ReportAllocs()
		
		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})
}