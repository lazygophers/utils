package candy

import (
	"math/rand"
	"strconv"
	"testing"
)

// 测试数据集
var (
	// Small dataset
	smallInts    = generateInts(10)
	smallStrings = generateStrings(10)
	smallFloats  = generateFloats(10)

	// Medium dataset
	mediumInts    = generateInts(100)
	mediumStrings = generateStrings(100)
	mediumFloats  = generateFloats(100)

	// Large dataset
	largeInts    = generateInts(1000)
	largeStrings = generateStrings(1000)
	largeFloats  = generateFloats(1000)

	// Huge dataset
	hugeInts    = generateInts(10000)
	hugeStrings = generateStrings(10000)
	hugeFloats  = generateFloats(10000)

	// Struct dataset
	smallPoints = generatePoints(10)
	largePoints = generatePoints(1000)
)

type point struct{ X, Y int }

func generateInts(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	return s
}

func generateStrings(n int) []string {
	s := make([]string, n)
	for i := range s {
		s[i] = strconv.Itoa(i)
	}
	return s
}

func generateFloats(n int) []float64 {
	s := make([]float64, n)
	for i := range s {
		s[i] = rand.Float64() * 100
	}
	return s
}

func generatePoints(n int) []point {
	p := make([]point, n)
	for i := range p {
		p[i] = point{X: rand.Intn(100), Y: rand.Intn(100)}
	}
	return p
}

// ============ 方案 1: 原始实现（基准） ============
func reduce1[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}
	result := ss[0]
	for _, s := range ss[1:] {
		result = f(result, s)
	}
	return result
}

// ============ 方案 2: 避免切片分配 ============
func reduce2[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}
	result := ss[0]
	for i := 1; i < len(ss); i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 3: 手动内联小数组 ============
func reduce3[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 4: 循环展开 x2 ============
func reduce4[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	result := ss[0]
	i := 1
	for ; i+1 < n; i += 2 {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
	}
	for ; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 5: 循环展开 x4 ============
func reduce5[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	result := ss[0]
	i := 1
	for ; i+3 < n; i += 4 {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
	}
	for ; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 6: 循环展开 x8 ============
func reduce6[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	result := ss[0]
	i := 1
	for ; i+7 < n; i += 8 {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
		result = f(result, ss[i+4])
		result = f(result, ss[i+5])
		result = f(result, ss[i+6])
		result = f(result, ss[i+7])
	}
	for ; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 7: 针对小数据集优化 ============
func reduce7[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	if n == 2 {
		return f(ss[0], ss[1])
	}
	if n == 3 {
		return f(f(ss[0], ss[1]), ss[2])
	}
	if n == 4 {
		return f(f(f(ss[0], ss[1]), ss[2]), ss[3])
	}
	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 8: 消除边界检查 ============
func reduce8[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 9: 使用 while 风格循环 ============
func reduce9[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	result := ss[0]
	i := 1
	for i < n {
		result = f(result, ss[i])
		i++
	}
	return result
}

// ============ 方案 10: 分治策略 ============
func reduce10[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	if n <= 32 {
		result := ss[0]
		for i := 1; i < n; i++ {
			result = f(result, ss[i])
		}
		return result
	}
	mid := n / 2
	left := reduce10(ss[:mid], f)
	right := reduce10(ss[mid:], f)
	return f(left, right)
}

// ============ 方案 11: 混合策略（小数据集特化 + 展开） ============
func reduce11[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}
	if n == 2 {
		return f(ss[0], ss[1])
	}
	if n == 3 {
		return f(f(ss[0], ss[1]), ss[2])
	}
	result := ss[0]
	i := 1
	for ; i+3 < n; i += 4 {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
	}
	for ; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

// ============ 方案 12: 针对不同大小优化 ============
func reduce12[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	result := ss[0]
	switch {
	case n <= 4:
		for i := 1; i < n; i++ {
			result = f(result, ss[i])
		}
	case n <= 16:
		for i := 1; i < n; i++ {
			result = f(result, ss[i])
		}
	default:
		i := 1
		for ; i+7 < n; i += 8 {
			result = f(result, ss[i])
			result = f(result, ss[i+1])
			result = f(result, ss[i+2])
			result = f(result, ss[i+3])
			result = f(result, ss[i+4])
			result = f(result, ss[i+5])
			result = f(result, ss[i+6])
			result = f(result, ss[i+7])
		}
		for ; i < n; i++ {
			result = f(result, ss[i])
		}
	}
	return result
}

// ============ 基准测试 ============

// 求和函数
func sumInt(a, b int) int { return a + b }
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func prodInt(a, b int) int            { return a * b }
func concatString(a, b string) string { return a + b }
func sumFloat(a, b float64) float64   { return a + b }
func sumPoint(a, b point) point       { return point{X: a.X + b.X, Y: a.Y + b.Y} }

// 方案1: 原始实现
func BenchmarkReduce1_SmallInt_Sum(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = reduce1(smallInts, sumInt)
	}
	_ = result
}

func BenchmarkReduce1_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(mediumInts, sumInt)
	}
}

func BenchmarkReduce1_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(largeInts, sumInt)
	}
}

func BenchmarkReduce1_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(hugeInts, sumInt)
	}
}

func BenchmarkReduce1_LargeInt_Max(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(largeInts, maxInt)
	}
}

func BenchmarkReduce1_LargeString_Concat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(largeStrings, concatString)
	}
}

func BenchmarkReduce1_LargeFloat_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(largeFloats, sumFloat)
	}
}

func BenchmarkReduce1_LargePoint_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce1(largePoints, sumPoint)
	}
}

// 方案2: 避免切片分配
func BenchmarkReduce2_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce2(smallInts, sumInt)
	}
}

func BenchmarkReduce2_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce2(mediumInts, sumInt)
	}
}

func BenchmarkReduce2_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce2(largeInts, sumInt)
	}
}

func BenchmarkReduce2_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce2(hugeInts, sumInt)
	}
}

// 方案3: 手动内联小数组
func BenchmarkReduce3_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce3(smallInts, sumInt)
	}
}

func BenchmarkReduce3_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce3(mediumInts, sumInt)
	}
}

func BenchmarkReduce3_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce3(largeInts, sumInt)
	}
}

func BenchmarkReduce3_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce3(hugeInts, sumInt)
	}
}

// 方案4: 循环展开 x2
func BenchmarkReduce4_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce4(smallInts, sumInt)
	}
}

func BenchmarkReduce4_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce4(mediumInts, sumInt)
	}
}

func BenchmarkReduce4_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce4(largeInts, sumInt)
	}
}

func BenchmarkReduce4_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce4(hugeInts, sumInt)
	}
}

// 方案5: 循环展开 x4
func BenchmarkReduce5_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce5(smallInts, sumInt)
	}
}

func BenchmarkReduce5_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce5(mediumInts, sumInt)
	}
}

func BenchmarkReduce5_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce5(largeInts, sumInt)
	}
}

func BenchmarkReduce5_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce5(hugeInts, sumInt)
	}
}

// 方案6: 循环展开 x8
func BenchmarkReduce6_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce6(smallInts, sumInt)
	}
}

func BenchmarkReduce6_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce6(mediumInts, sumInt)
	}
}

func BenchmarkReduce6_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce6(largeInts, sumInt)
	}
}

func BenchmarkReduce6_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce6(hugeInts, sumInt)
	}
}

// 方案7: 针对小数据集优化
func BenchmarkReduce7_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce7(smallInts, sumInt)
	}
}

func BenchmarkReduce7_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce7(mediumInts, sumInt)
	}
}

func BenchmarkReduce7_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce7(largeInts, sumInt)
	}
}

func BenchmarkReduce7_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce7(hugeInts, sumInt)
	}
}

// 方案8: 消除边界检查
func BenchmarkReduce8_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce8(smallInts, sumInt)
	}
}

func BenchmarkReduce8_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce8(mediumInts, sumInt)
	}
}

func BenchmarkReduce8_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce8(largeInts, sumInt)
	}
}

func BenchmarkReduce8_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce8(hugeInts, sumInt)
	}
}

// 方案9: 使用 while 风格循环
func BenchmarkReduce9_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce9(smallInts, sumInt)
	}
}

func BenchmarkReduce9_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce9(mediumInts, sumInt)
	}
}

func BenchmarkReduce9_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce9(largeInts, sumInt)
	}
}

func BenchmarkReduce9_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce9(hugeInts, sumInt)
	}
}

// 方案10: 分治策略
func BenchmarkReduce10_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce10(smallInts, sumInt)
	}
}

func BenchmarkReduce10_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce10(mediumInts, sumInt)
	}
}

func BenchmarkReduce10_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce10(largeInts, sumInt)
	}
}

func BenchmarkReduce10_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce10(hugeInts, sumInt)
	}
}

// 方案11: 混合策略
func BenchmarkReduce11_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce11(smallInts, sumInt)
	}
}

func BenchmarkReduce11_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce11(mediumInts, sumInt)
	}
}

func BenchmarkReduce11_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce11(largeInts, sumInt)
	}
}

func BenchmarkReduce11_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce11(hugeInts, sumInt)
	}
}

// 方案12: 针对不同大小优化
func BenchmarkReduce12_SmallInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce12(smallInts, sumInt)
	}
}

func BenchmarkReduce12_MediumInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce12(mediumInts, sumInt)
	}
}

func BenchmarkReduce12_LargeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce12(largeInts, sumInt)
	}
}

func BenchmarkReduce12_HugeInt_Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reduce12(hugeInts, sumInt)
	}
}
