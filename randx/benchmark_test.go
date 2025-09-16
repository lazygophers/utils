package randx

import (
	"math/rand"
	"testing"
	"time"
)

// 原始实现的函数用于对比（性能低）
func originalIntn(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

func originalInt64() int64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
}

func originalFloat64() float64 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
}

// 基准测试：随机整数生成
func BenchmarkIntn(b *testing.B) {
	b.Run("Original_Intn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = originalIntn(100)
		}
	})

	b.Run("Optimized_Intn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})

	b.Run("Fast_Intn", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})
}

// 基准测试：Int64生成
func BenchmarkInt64(b *testing.B) {
	b.Run("Original_Int64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = originalInt64()
		}
	})

	b.Run("Optimized_Int64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int64()
		}
	})

	b.Run("Fast_Int64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int()
		}
	})
}

// 基准测试：Float64生成
func BenchmarkFloat64(b *testing.B) {
	b.Run("Original_Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = originalFloat64()
		}
	})

	b.Run("Optimized_Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Float64()
		}
	})

	b.Run("Fast_Float64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Float64()
		}
	})
}

// 基准测试：范围随机数
func BenchmarkRange(b *testing.B) {
	b.Run("IntnRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = IntnRange(1, 100)
		}
	})

	b.Run("Int64nRange", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int64nRange(1, 100)
		}
	})

	b.Run("Float64Range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Float64Range(1.0, 100.0)
		}
	})
}

// 基准测试：布尔值生成
func BenchmarkBool(b *testing.B) {
	b.Run("Bool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Bool()
		}
	})

	b.Run("FastBool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Bool()
		}
	})

	b.Run("WeightedBool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = WeightedBool(0.5)
		}
	})
}

// 基准测试：选择函数
func BenchmarkChoose(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b.Run("Choose", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Choose(slice)
		}
	})

	b.Run("FastChoose", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Choose(slice)
		}
	})
}

// 基准测试：时间相关函数
func BenchmarkTime(b *testing.B) {
	b.Run("TimeDuration4Sleep", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = TimeDuration4Sleep(time.Millisecond, time.Second)
		}
	})

	b.Run("FastTimeDuration4Sleep", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = TimeDuration4Sleep(time.Millisecond, time.Second)
		}
	})

	b.Run("RandomDuration", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = RandomDuration(time.Millisecond, time.Second)
		}
	})
}

// 基准测试：批量操作
func BenchmarkBatch(b *testing.B) {
	b.Run("BatchIntn_Single", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})

	b.Run("BatchIntn_Batch10", func(b *testing.B) {
		for i := 0; i < b.N; i += 10 {
			_ = BatchIntn(100, 10)
		}
	})

	b.Run("BatchInt64n_Single", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Int64n(100)
		}
	})

	b.Run("BatchInt64n_Batch10", func(b *testing.B) {
		for i := 0; i < b.N; i += 10 {
			_ = BatchInt64n(100, 10)
		}
	})
}

// 基准测试：并发性能
func BenchmarkConcurrent(b *testing.B) {
	b.Run("Intn_Sequential", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})

	b.Run("Intn_Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Intn(100)
			}
		})
	})

	b.Run("FastIntn_Sequential", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})

	b.Run("FastIntn_Parallel", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = Intn(100)
			}
		})
	})
}

// 基准测试：内存分配
func BenchmarkMemoryAllocation(b *testing.B) {
	b.Run("Original_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = originalIntn(100)
		}
	})

	b.Run("Optimized_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})

	b.Run("Fast_Memory", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
		}
	})
}

// 基准测试：复杂操作
func BenchmarkComplexOperations(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.Run("Shuffle", func(b *testing.B) {
		b.StopTimer()
		testSlice := make([]int, len(slice))
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(testSlice, slice)
			b.StartTimer()
			Shuffle(testSlice)
		}
	})

	b.Run("FastShuffle", func(b *testing.B) {
		b.StopTimer()
		testSlice := make([]int, len(slice))
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(testSlice, slice)
			b.StartTimer()
			Shuffle(testSlice)
		}
	})

	b.Run("ChooseN", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ChooseN(slice, 10)
		}
	})
}

// 基准测试：高频调用场景
func BenchmarkHighFrequency(b *testing.B) {
	b.Run("MixedOperations_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = originalIntn(100)
			_ = originalFloat64()
			_ = originalInt64()
		}
	})

	b.Run("MixedOperations_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
			_ = Float64()
			_ = Int64()
		}
	})

	b.Run("MixedOperations_Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Intn(100)
			_ = Float64()
			_ = Int()
		}
	})
}
