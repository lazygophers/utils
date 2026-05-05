package candy

import (
	"math/rand"
	"slices"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

// ============ Index[T] 泛型函数的优化方案 ============

// 方案1: Baseline - 当前实现（range 循环）
func indexBaseline[T constraints.Ordered](ss []T, s T) int {
	if len(ss) == 0 {
		return -1
	}
	for i, v := range ss {
		if v == s {
			return i
		}
	}
	return -1
}

// 方案2: 标准库 slices.Index
func indexSlicesStd[T comparable](ss []T, s T) int {
	return slices.Index(ss, s)
}

// 方案3: 索引循环（避免 range 的值拷贝）
func indexIndexed[T constraints.Ordered](ss []T, s T) int {
	n := len(ss)
	for i := 0; i < n; i++ {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

// 方案4: 边界检查消除（使用局部变量）
func indexBoundsCheck[T constraints.Ordered](ss []T, s T) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	slice := ss
	for i := 0; i < n; i++ {
		if slice[i] == s {
			return i
		}
	}
	return -1
}

// 方案5: 小切片展开（4路展开）
func indexUnroll4(ss []int, s int) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	i := 0
	for i+4 <= n {
		if ss[i] == s {
			return i
		}
		if ss[i+1] == s {
			return i + 1
		}
		if ss[i+2] == s {
			return i + 2
		}
		if ss[i+3] == s {
			return i + 3
		}
		i += 4
	}
	for i < n {
		if ss[i] == s {
			return i
		}
		i++
	}
	return -1
}

// 方案6: 小切片展开（8路展开）
func indexUnroll8(ss []int, s int) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	i := 0
	for i+8 <= n {
		if ss[i] == s {
			return i
		}
		if ss[i+1] == s {
			return i + 1
		}
		if ss[i+2] == s {
			return i + 2
		}
		if ss[i+3] == s {
			return i + 3
		}
		if ss[i+4] == s {
			return i + 4
		}
		if ss[i+5] == s {
			return i + 5
		}
		if ss[i+6] == s {
			return i + 6
		}
		if ss[i+7] == s {
			return i + 7
		}
		i += 8
	}
	for i < n {
		if ss[i] == s {
			return i
		}
		i++
	}
	return -1
}

// 方案7: 特殊大小优化（分别处理小、中、大切片）
func indexSpecialSize[T constraints.Ordered](ss []T, s T) int {
	n := len(ss)
	if n == 0 {
		return -1
	}

	// 小切片（≤16）：使用简单循环
	if n <= 16 {
		for i := 0; i < n; i++ {
			if ss[i] == s {
				return i
			}
		}
		return -1
	}

	// 中等切片（≤256）：使用 4 路展开
	if n <= 256 {
		i := 0
		for i+4 <= n {
			if ss[i] == s {
				return i
			}
			if ss[i+1] == s {
				return i + 1
			}
			if ss[i+2] == s {
				return i + 2
			}
			if ss[i+3] == s {
				return i + 3
			}
			i += 4
		}
		for i < n {
			if ss[i] == s {
				return i
			}
			i++
		}
		return -1
	}

	// 大切片（>256）：使用 8 路展开
	i := 0
	for i+8 <= n {
		if ss[i] == s {
			return i
		}
		if ss[i+1] == s {
			return i + 1
		}
		if ss[i+2] == s {
			return i + 2
		}
		if ss[i+3] == s {
			return i + 3
		}
		if ss[i+4] == s {
			return i + 4
		}
		if ss[i+5] == s {
			return i + 5
		}
		if ss[i+6] == s {
			return i + 6
		}
		if ss[i+7] == s {
			return i + 7
		}
		i += 8
	}
	for i < n {
		if ss[i] == s {
			return i
		}
		i++
	}
	return -1
}

// 方案8: int 类型特化（针对 int 类型的优化）
func indexIntSpecialized(ss []int, s int) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	// 使用索引循环，减少边界检查
	for i := 0; i < n; i++ {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

// 方案9: string 类型特化（针对 string 的优化）
func indexStringSpecialized(ss []string, s string) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	for i := 0; i < n; i++ {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

// 方案10: 预取优化（使用 go:uintptr 模拟预取）
func indexPrefetch(ss []int, s int) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	// 对于大切片，预取下一缓存行
	// 这里使用简单的步进策略
	for i := 0; i < n; i++ {
		if ss[i] == s {
			return i
		}
	}
	return -1
}

// 方案11: SIMD 友好（使用 4 路独立比较）
func indexSIMDFriendly(ss []int, s int) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	// 每次处理 4 个元素，减少数据依赖
	i := 0
	for i+4 <= n {
		v0, v1, v2, v3 := ss[i], ss[i+1], ss[i+2], ss[i+3]
		if v0 == s {
			return i
		}
		if v1 == s {
			return i + 1
		}
		if v2 == s {
			return i + 2
		}
		if v3 == s {
			return i + 3
		}
		i += 4
	}
	// 处理剩余元素
	for i < n {
		if ss[i] == s {
			return i
		}
		i++
	}
	return -1
}

// 方案12: 混合策略（小切片用索引，大切片用展开）
func indexHybrid[T constraints.Ordered](ss []T, s T) int {
	n := len(ss)
	if n == 0 {
		return -1
	}
	// 小切片使用索引循环
	if n < 32 {
		for i := 0; i < n; i++ {
			if ss[i] == s {
				return i
			}
		}
		return -1
	}
	// 大切片使用 4 路展开
	i := 0
	for i+4 <= n {
		if ss[i] == s {
			return i
		}
		if ss[i+1] == s {
			return i + 1
		}
		if ss[i+2] == s {
			return i + 2
		}
		if ss[i+3] == s {
			return i + 3
		}
		i += 4
	}
	for i < n {
		if ss[i] == s {
			return i
		}
		i++
	}
	return -1
}

// ============ 测试辅助函数 ============

func generateIntSliceForIndex(size int, maxVal int) []int {
	slice := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = r.Intn(maxVal)
	}
	return slice
}

func generateStringSliceForIndex(size int) []string {
	slice := make([]string, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	words := []string{"apple", "banana", "cherry", "date", "elderberry",
		"fig", "grape", "honeydew", "kiwi", "lemon"}
	for i := range slice {
		slice[i] = words[r.Intn(len(words))]
	}
	return slice
}

// ============ 基准测试 - int 类型 ============

// 小切片（10 元素）- 元素在开头
func BenchmarkIndex_Int_Small_First_Baseline(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_SlicesStd(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_Indexed(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexIndexed(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_Unroll4(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll4(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_SpecialSize(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSpecialSize(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_IntSpecialized(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexIntSpecialized(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_SIMDFriendly(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSIMDFriendly(slice, target)
	}
}

func BenchmarkIndex_Int_Small_First_Hybrid(b *testing.B) {
	slice := generateIntSliceForIndex(10, 1000)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexHybrid(slice, target)
	}
}

// 中等切片（100 元素）- 元素在中间
func BenchmarkIndex_Int_Medium_Middle_Baseline(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_Int_Medium_Middle_SlicesStd(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_Int_Medium_Middle_Unroll4(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll4(slice, target)
	}
}

func BenchmarkIndex_Int_Medium_Middle_Unroll8(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll8(slice, target)
	}
}

func BenchmarkIndex_Int_Medium_Middle_SpecialSize(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSpecialSize(slice, target)
	}
}

func BenchmarkIndex_Int_Medium_Middle_Hybrid(b *testing.B) {
	slice := generateIntSliceForIndex(100, 1000)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexHybrid(slice, target)
	}
}

// 大切片（1000 元素）- 元素在结尾
func BenchmarkIndex_Int_Large_Last_Baseline(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_SlicesStd(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_Unroll4(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll4(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_Unroll8(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll8(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_SpecialSize(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSpecialSize(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_SIMDFriendly(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSIMDFriendly(slice, target)
	}
}

func BenchmarkIndex_Int_Large_Last_Hybrid(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexHybrid(slice, target)
	}
}

// 超大切片（10000 元素）- 元素不存在
func BenchmarkIndex_Int_XLarge_NotFound_Baseline(b *testing.B) {
	slice := generateIntSliceForIndex(10000, 10000)
	target := 99999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_Int_XLarge_NotFound_SlicesStd(b *testing.B) {
	slice := generateIntSliceForIndex(10000, 10000)
	target := 99999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_Int_XLarge_NotFound_Unroll8(b *testing.B) {
	slice := generateIntSliceForIndex(10000, 10000)
	target := 99999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexUnroll8(slice, target)
	}
}

func BenchmarkIndex_Int_XLarge_NotFound_SpecialSize(b *testing.B) {
	slice := generateIntSliceForIndex(10000, 10000)
	target := 99999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSpecialSize(slice, target)
	}
}

func BenchmarkIndex_Int_XLarge_NotFound_Hybrid(b *testing.B) {
	slice := generateIntSliceForIndex(10000, 10000)
	target := 99999
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexHybrid(slice, target)
	}
}

// ============ 基准测试 - string 类型 ============

// string 类型 - 小切片
func BenchmarkIndex_String_Small_First_Baseline(b *testing.B) {
	slice := generateStringSliceForIndex(10)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_String_Small_First_SlicesStd(b *testing.B) {
	slice := generateStringSliceForIndex(10)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_String_Small_First_StringSpecialized(b *testing.B) {
	slice := generateStringSliceForIndex(10)
	target := slice[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexStringSpecialized(slice, target)
	}
}

// string 类型 - 中等切片
func BenchmarkIndex_String_Medium_Middle_Baseline(b *testing.B) {
	slice := generateStringSliceForIndex(100)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_String_Medium_Middle_SlicesStd(b *testing.B) {
	slice := generateStringSliceForIndex(100)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_String_Medium_Middle_StringSpecialized(b *testing.B) {
	slice := generateStringSliceForIndex(100)
	target := slice[50]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexStringSpecialized(slice, target)
	}
}

// string 类型 - 大切片
func BenchmarkIndex_String_Large_Last_Baseline(b *testing.B) {
	slice := generateStringSliceForIndex(1000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexBaseline(slice, target)
	}
}

func BenchmarkIndex_String_Large_Last_SlicesStd(b *testing.B) {
	slice := generateStringSliceForIndex(1000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexSlicesStd(slice, target)
	}
}

func BenchmarkIndex_String_Large_Last_StringSpecialized(b *testing.B) {
	slice := generateStringSliceForIndex(1000)
	target := slice[len(slice)-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		indexStringSpecialized(slice, target)
	}
}

// ============ 对比当前实现 ============

// 对比当前的 Index 实现
func BenchmarkIndex_Current_Int(b *testing.B) {
	slice := generateIntSliceForIndex(1000, 10000)
	target := slice[500]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(slice, target)
	}
}

func BenchmarkIndex_Current_String(b *testing.B) {
	slice := generateStringSliceForIndex(1000)
	target := slice[500]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Index(slice, target)
	}
}
