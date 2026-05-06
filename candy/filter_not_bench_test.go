package candy

import "testing"

// 准备测试数据
var (
	smallDataInt   []int
	mediumDataInt  []int
	largeDataInt   []int
	smallDataStr   []string
	mediumDataStr  []string
	largeDataStr   []string
)

func init() {
	// 小数据集: 100 个元素
	smallDataInt = make([]int, 100)
	smallDataStr = make([]string, 100)
	for i := 0; i < 100; i++ {
		smallDataInt[i] = i
		smallDataStr[i] = string(rune('a' + i%26))
	}

	// 中等数据集: 1000 个元素
	mediumDataInt = make([]int, 1000)
	mediumDataStr = make([]string, 1000)
	for i := 0; i < 1000; i++ {
		mediumDataInt[i] = i
		mediumDataStr[i] = string(rune('a' + i%26))
	}

	// 大数据集: 10000 个元素
	largeDataInt = make([]int, 10000)
	largeDataStr = make([]string, 10000)
	for i := 0; i < 10000; i++ {
		largeDataInt[i] = i
		largeDataStr[i] = string(rune('a' + i%26))
	}
}

// 基准测试 - 原始实现
func BenchmarkFilterNot_Original_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_Original_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_Original_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FilterNot(largeDataInt, predicate)
	}
}

// 方案1: 预分配一半容量
func BenchmarkFilterNot_V1_PreHalf_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV1_PreHalf(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V1_PreHalf_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV1_PreHalf(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V1_PreHalf_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV1_PreHalf(largeDataInt, predicate)
	}
}

// 方案2: 预分配全容量
func BenchmarkFilterNot_V2_PreFull_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV2_PreFull(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V2_PreFull_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV2_PreFull(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V2_PreFull_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV2_PreFull(largeDataInt, predicate)
	}
}

// 方案3: 索引循环
func BenchmarkFilterNot_V3_IndexLoop_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV3_IndexLoop(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V3_IndexLoop_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV3_IndexLoop(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V3_IndexLoop_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV3_IndexLoop(largeDataInt, predicate)
	}
}

// 方案4: 预分配一半 + 索引循环
func BenchmarkFilterNot_V4_PreHalfIndex_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV4_PreHalfIndex(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V4_PreHalfIndex_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV4_PreHalfIndex(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V4_PreHalfIndex_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV4_PreHalfIndex(largeDataInt, predicate)
	}
}

// 方案5: 预分配全容量 + 索引循环
func BenchmarkFilterNot_V5_PreFullIndex_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV5_PreFullIndex(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V5_PreFullIndex_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV5_PreFullIndex(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V5_PreFullIndex_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV5_PreFullIndex(largeDataInt, predicate)
	}
}

// 方案6: 两遍扫描优化
func BenchmarkFilterNot_V6_TwoPass_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV6_TwoPass(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V6_TwoPass_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV6_TwoPass(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V6_TwoPass_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV6_TwoPass(largeDataInt, predicate)
	}
}

// 方案7: 复制后切片
func BenchmarkFilterNot_V7_CopySlice_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV7_CopySlice(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V7_CopySlice_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV7_CopySlice(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V7_CopySlice_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV7_CopySlice(largeDataInt, predicate)
	}
}

// 方案8: 原地修改（非零拷贝）
func BenchmarkFilterNot_V8_InPlace_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV8_InPlace(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V8_InPlace_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV8_InPlace(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V8_InPlace_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV8_InPlace(largeDataInt, predicate)
	}
}

// 方案9: 动态容量调整
func BenchmarkFilterNot_V9_Dynamic_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV9_Dynamic(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V9_Dynamic_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV9_Dynamic(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V9_Dynamic_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV9_Dynamic(largeDataInt, predicate)
	}
}

// 方案10: 空切片快速路径
func BenchmarkFilterNot_V10_FastPath_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV10_FastPath(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V10_FastPath_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV10_FastPath(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V10_FastPath_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV10_FastPath(largeDataInt, predicate)
	}
}

// 方案11: Filter + Not 逻辑
func BenchmarkFilterNot_V11_Composed_Small(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV11_Composed(smallDataInt, predicate)
	}
}

func BenchmarkFilterNot_V11_Composed_Medium(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV11_Composed(mediumDataInt, predicate)
	}
}

func BenchmarkFilterNot_V11_Composed_Large(b *testing.B) {
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filterNotV11_Composed(largeDataInt, predicate)
	}
}
