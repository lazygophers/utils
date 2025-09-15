package fake

import (
	"sync/atomic"
	"testing"
)

// 纳秒级性能对比测试
func BenchmarkNanoComparison(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		faker := New()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.Name()
		}
	})

	b.Run("PreviousExtreme", func(b *testing.B) {
		faker := NewExtremePerformance()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ExtremeName()
		}
	})

	b.Run("Nano", func(b *testing.B) {
		faker := NewNanoPerformance()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.NanoName()
		}
	})

	b.Run("Atomic", func(b *testing.B) {
		faker := NewAtomic()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.AtomicName()
		}
	})

	b.Run("Constant", func(b *testing.B) {
		faker := NewConstant()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.ConstantName()
		}
	})

	b.Run("IncrementOnly", func(b *testing.B) {
		faker := NewIncrementOnly()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.IncrementOnlyName()
		}
	})

	b.Run("Static", func(b *testing.B) {
		faker := NewStatic()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.StaticName()
		}
	})

	b.Run("CPUOptimized", func(b *testing.B) {
		faker := NewCPUOptimized()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.CPUOptimizedName()
		}
	})

	b.Run("Branchless", func(b *testing.B) {
		faker := NewBranchless()
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = faker.BranchlessName()
		}
	})

	b.Run("Ultimate", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = UltimatePerformanceName()
		}
	})

	// 全局函数测试
	b.Run("GlobalNano", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = NanoName()
		}
	})

	b.Run("GlobalConstant", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = ConstantName()
		}
	})

	b.Run("GlobalStatic", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = StaticName()
		}
	})
}

// 原子操作性能测试
func BenchmarkAtomicOperationComparison(b *testing.B) {
	b.Run("AtomicAdd", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&counter, 1)
		}
	})

	b.Run("AtomicLoad", func(b *testing.B) {
		var counter uint64 = 1
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = atomic.LoadUint64(&counter)
		}
	})

	b.Run("AtomicStore", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			atomic.StoreUint64(&counter, uint64(i))
		}
	})

	b.Run("AtomicCAS", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			atomic.CompareAndSwapUint64(&counter, 0, 1)
			atomic.CompareAndSwapUint64(&counter, 1, 0)
		}
	})

	b.Run("FastAtomicAdd", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = FastAtomicAdd()
		}
	})

	b.Run("FastAtomicLoad", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = FastAtomicLoad()
		}
	})

	b.Run("FastAtomicCAS", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = FastAtomicCAS()
		}
	})
}

// 并发纳秒级性能测试
func BenchmarkNanoConcurrent(b *testing.B) {
	b.Run("NanoConcurrent", func(b *testing.B) {
		faker := NewNanoPerformance()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.NanoName()
			}
		})
	})

	b.Run("ConstantConcurrent", func(b *testing.B) {
		faker := NewConstant()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.ConstantName()
			}
		})
	})

	b.Run("StaticConcurrent", func(b *testing.B) {
		faker := NewStatic()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.StaticName()
			}
		})
	})

	b.Run("UltimateConcurrent", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = UltimatePerformanceName()
			}
		})
	})

	b.Run("BranchlessConcurrent", func(b *testing.B) {
		faker := NewBranchless()
		b.ResetTimer()
		b.ReportAllocs()

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = faker.BranchlessName()
			}
		})
	})
}

// CPU指令级性能测试
func BenchmarkCPUInstructions(b *testing.B) {
	b.Run("SimpleReturn", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = "John Smith" // 编译器常量
		}
	})

	b.Run("FunctionCall", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = UltimatePerformanceName()
		}
	})

	b.Run("AtomicIncrement", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&counter, 1)
		}
	})

	b.Run("BitOperation", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			counter++
			_ = counter & 3 // 位运算
		}
	})

	b.Run("ArrayAccess", func(b *testing.B) {
		names := [4]string{"A", "B", "C", "D"}
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			counter++
			_ = names[counter&3]
		}
	})

	b.Run("SwitchStatement", func(b *testing.B) {
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			counter++
			idx := counter & 3
			var result string
			switch idx {
			case 0:
				result = "A"
			case 1:
				result = "B"
			case 2:
				result = "C"
			default:
				result = "D"
			}
			_ = result
		}
	})
}

// 内存访问模式测试
func BenchmarkMemoryPatterns(b *testing.B) {
	b.Run("StackVariable", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			name := "John Smith"
			_ = name
		}
	})

	b.Run("GlobalConstant", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = StaticName()
		}
	})

	b.Run("HeapAllocation", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			name := new(string)
			*name = "John Smith"
			_ = *name
		}
	})

	b.Run("StringConcatenation", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_ = "John" + " " + "Smith"
		}
	})
}

// 测试编译器优化影响
func BenchmarkCompilerOptimizations(b *testing.B) {
	b.Run("InlineConstant", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			_ = "John Smith" // 编译器会内联
		}
	})

	b.Run("NonInlineFunction", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// 使用函数调用避免内联
			_ = func() string { return "John Smith" }()
		}
	})

	b.Run("VolatileAccess", func(b *testing.B) {
		names := [4]string{"John Smith", "Mary Johnson", "James Williams", "Patricia Brown"}
		var counter uint64
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			// 模拟无法内联优化的访问
			idx := atomic.AddUint64(&counter, 1) & 3
			_ = names[idx]
		}
	})
}
