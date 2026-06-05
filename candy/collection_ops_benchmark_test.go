package candy

import (
	stdrand "math/rand"
	"math/rand/v2"
	"slices"
	"sort"
	"strconv"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

// ==================== 测试数据生成 ====================

// 生成整数切片
func genIntsPred(n int) []int {
	r := stdrand.New(stdrand.NewSource(42))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Int()
	}
	return s
}

// ==================== 回调函数 ====================

// All/Any 的谓词函数
func isPositive(v int) bool {
	return v > 0
}

// ==================== All 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）- 索引循环
func All1[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			return false
		}
	}
	return true
}

// 方案2: 使用 range 循环
func All2[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if !f(s) {
			return false
		}
	}
	return true
}

// 方案3: 空切片快速返回
func All3[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return true
	}
	n := len(ss)
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			return false
		}
	}
	return true
}

// 方案4: 使用 range + 空切片检查
func All4[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return true
	}
	for _, s := range ss {
		if !f(s) {
			return false
		}
	}
	return true
}

// 方案5: 循环展开（4路）
func All5[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0

	for i < n && (n-i)%4 != 0 {
		if !f(ss[i]) {
			return false
		}
		i++
	}

	for i+3 < n {
		if !f(ss[i]) || !f(ss[i+1]) || !f(ss[i+2]) || !f(ss[i+3]) {
			return false
		}
		i += 4
	}

	for i < n {
		if !f(ss[i]) {
			return false
		}
		i++
	}
	return true
}

// 方案6: 循环展开（8路）
func All6[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0

	for i < n && (n-i)%8 != 0 {
		if !f(ss[i]) {
			return false
		}
		i++
	}

	for i+7 < n {
		if !f(ss[i]) || !f(ss[i+1]) || !f(ss[i+2]) || !f(ss[i+3]) ||
			!f(ss[i+4]) || !f(ss[i+5]) || !f(ss[i+6]) || !f(ss[i+7]) {
			return false
		}
		i += 8
	}

	for i < n {
		if !f(ss[i]) {
			return false
		}
		i++
	}
	return true
}

// 方案7: while 风格
func All7[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0
	for i < n {
		if !f(ss[i]) {
			return false
		}
		i++
	}
	return true
}

// 方案8: 特殊情况处理
func All8[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	if n == 0 {
		return true
	}
	if n == 1 {
		return f(ss[0])
	}
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			return false
		}
	}
	return true
}

// 方案9: 局部变量优化
func All9[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	var i int
	for i = 0; i < n; i++ {
		if !f(ss[i]) {
			return false
		}
	}
	return true
}

// 方案10: 混合策略
func All10[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	if n <= 1 {
		if n == 1 {
			return f(ss[0])
		}
		return true
	}
	for i := 0; i < n; i++ {
		if !f(ss[i]) {
			return false
		}
	}
	return true
}

// ==================== Any 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）- 索引循环
func Any1[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	for i := 0; i < n; i++ {
		if f(ss[i]) {
			return true
		}
	}
	return false
}

// 方案2: 使用 range 循环
func Any2[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}
	return false
}

// 方案3: 空切片快速返回
func Any3[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return false
	}
	n := len(ss)
	for i := 0; i < n; i++ {
		if f(ss[i]) {
			return true
		}
	}
	return false
}

// 方案4: 使用 range + 空切片检查
func Any4[T any](ss []T, f func(T) bool) bool {
	if len(ss) == 0 {
		return false
	}
	for _, s := range ss {
		if f(s) {
			return true
		}
	}
	return false
}

// 方案5: 循环展开（4路）
func Any5[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0

	for i < n && (n-i)%4 != 0 {
		if f(ss[i]) {
			return true
		}
		i++
	}

	for i+3 < n {
		if f(ss[i]) || f(ss[i+1]) || f(ss[i+2]) || f(ss[i+3]) {
			return true
		}
		i += 4
	}

	for i < n {
		if f(ss[i]) {
			return true
		}
		i++
	}
	return false
}

// 方案6: 循环展开（8路）
func Any6[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0

	for i < n && (n-i)%8 != 0 {
		if f(ss[i]) {
			return true
		}
		i++
	}

	for i+7 < n {
		if f(ss[i]) || f(ss[i+1]) || f(ss[i+2]) || f(ss[i+3]) ||
			f(ss[i+4]) || f(ss[i+5]) || f(ss[i+6]) || f(ss[i+7]) {
			return true
		}
		i += 8
	}

	for i < n {
		if f(ss[i]) {
			return true
		}
		i++
	}
	return false
}

// 方案7: while 风格
func Any7[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	i := 0
	for i < n {
		if f(ss[i]) {
			return true
		}
		i++
	}
	return false
}

// 方案8: 特殊情况处理
func Any8[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	if n == 0 {
		return false
	}
	if n == 1 {
		return f(ss[0])
	}
	for i := 0; i < n; i++ {
		if f(ss[i]) {
			return true
		}
	}
	return false
}

// 方案9: 局部变量优化
func Any9[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	var i int
	for i = 0; i < n; i++ {
		if f(ss[i]) {
			return true
		}
	}
	return false
}

// 方案10: 混合策略
func Any10[T any](ss []T, f func(T) bool) bool {
	n := len(ss)
	if n == 0 {
		return false
	}
	if n == 1 {
		return f(ss[0])
	}
	for i := 0; i < n; i++ {
		if f(ss[i]) {
			return true
		}
	}
	return false
}

// ==================== Benchmark 测试 ====================

// 小切片（10个元素）
func BenchmarkAll1_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All1(data, isPositive)
	}
}

func BenchmarkAll2_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All2(data, isPositive)
	}
}

func BenchmarkAll3_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All3(data, isPositive)
	}
}

func BenchmarkAll4_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All4(data, isPositive)
	}
}

func BenchmarkAll5_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All5(data, isPositive)
	}
}

func BenchmarkAll6_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All6(data, isPositive)
	}
}

func BenchmarkAll7_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All7(data, isPositive)
	}
}

func BenchmarkAll8_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All8(data, isPositive)
	}
}

func BenchmarkAll9_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All9(data, isPositive)
	}
}

func BenchmarkAll10_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All10(data, isPositive)
	}
}

// 中等切片（100个元素）
func BenchmarkAll1_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All1(data, isPositive)
	}
}

func BenchmarkAll2_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All2(data, isPositive)
	}
}

func BenchmarkAll3_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All3(data, isPositive)
	}
}

func BenchmarkAll4_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All4(data, isPositive)
	}
}

func BenchmarkAll5_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All5(data, isPositive)
	}
}

func BenchmarkAll6_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All6(data, isPositive)
	}
}

func BenchmarkAll7_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All7(data, isPositive)
	}
}

func BenchmarkAll8_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All8(data, isPositive)
	}
}

func BenchmarkAll9_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All9(data, isPositive)
	}
}

func BenchmarkAll10_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All10(data, isPositive)
	}
}

// 大切片（1000个元素）
func BenchmarkAll1_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All1(data, isPositive)
	}
}

func BenchmarkAll2_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All2(data, isPositive)
	}
}

func BenchmarkAll3_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All3(data, isPositive)
	}
}

func BenchmarkAll4_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All4(data, isPositive)
	}
}

func BenchmarkAll5_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All5(data, isPositive)
	}
}

func BenchmarkAll6_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All6(data, isPositive)
	}
}

func BenchmarkAll7_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All7(data, isPositive)
	}
}

func BenchmarkAll8_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All8(data, isPositive)
	}
}

func BenchmarkAll9_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All9(data, isPositive)
	}
}

func BenchmarkAll10_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		All10(data, isPositive)
	}
}

// ==================== Any Benchmarks ====================

// 小切片
func BenchmarkAny1_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any1(data, isPositive)
	}
}

func BenchmarkAny2_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any2(data, isPositive)
	}
}

func BenchmarkAny3_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any3(data, isPositive)
	}
}

func BenchmarkAny4_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any4(data, isPositive)
	}
}

func BenchmarkAny5_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any5(data, isPositive)
	}
}

func BenchmarkAny6_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any6(data, isPositive)
	}
}

func BenchmarkAny7_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any7(data, isPositive)
	}
}

func BenchmarkAny8_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any8(data, isPositive)
	}
}

func BenchmarkAny9_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any9(data, isPositive)
	}
}

func BenchmarkAny10_Small(b *testing.B) {
	data := genIntsPred(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any10(data, isPositive)
	}
}

// 中等切片
func BenchmarkAny1_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any1(data, isPositive)
	}
}

func BenchmarkAny2_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any2(data, isPositive)
	}
}

func BenchmarkAny3_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any3(data, isPositive)
	}
}

func BenchmarkAny4_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any4(data, isPositive)
	}
}

func BenchmarkAny5_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any5(data, isPositive)
	}
}

func BenchmarkAny6_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any6(data, isPositive)
	}
}

func BenchmarkAny7_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any7(data, isPositive)
	}
}

func BenchmarkAny8_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any8(data, isPositive)
	}
}

func BenchmarkAny9_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any9(data, isPositive)
	}
}

func BenchmarkAny10_Medium(b *testing.B) {
	data := genIntsPred(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any10(data, isPositive)
	}
}

// 大切片
func BenchmarkAny1_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any1(data, isPositive)
	}
}

func BenchmarkAny2_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any2(data, isPositive)
	}
}

func BenchmarkAny3_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any3(data, isPositive)
	}
}

func BenchmarkAny4_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any4(data, isPositive)
	}
}

func BenchmarkAny5_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any5(data, isPositive)
	}
}

func BenchmarkAny6_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any6(data, isPositive)
	}
}

func BenchmarkAny7_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any7(data, isPositive)
	}
}

func BenchmarkAny8_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any8(data, isPositive)
	}
}

func BenchmarkAny9_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any9(data, isPositive)
	}
}

func BenchmarkAny10_Large(b *testing.B) {
	data := genIntsPred(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Any10(data, isPositive)
	}
}

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
	r := stdrand.New(stdrand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		slice[i] = r.Intn(maxVal)
	}
	return slice
}

func generateStringSliceForIndex(size int) []string {
	slice := make([]string, size)
	r := stdrand.New(stdrand.NewSource(time.Now().UnixNano()))
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

// 生成测试数据
func genInts(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = int(rand.Uint64())
	}
	return s
}

func genStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa(int(rand.Uint64() % 1000))
	}
	return s
}

// 简单累加函数
func simpleAdd(a, b int) int {
	return a + b
}

// 复杂函数（模拟真实场景）
func complexOp(a, b int) int {
	time.Sleep(1 * time.Microsecond)
	return a + b + (a*b)/2
}

// 字符串连接
func concatStr(a, b string) string {
	return a + b
}

// ==================== 基准测试：不同大小的切片 ====================

// Baseline: 当前实现
func BenchmarkReduce_Current_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce_Current_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce_Current_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

// ==================== 10种优化方案 ====================

// 方案1: 避免切片重新分配，直接索引访问
func Reduce1[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}
	result := ss[0]
	for i := 1; i < len(ss); i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce1_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

func BenchmarkReduce1_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

// 方案2: 预检查并针对小切片优化
func Reduce2[T any](ss []T, f func(T, T) T) T {
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

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce2_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce2(data, simpleAdd)
	}
}

func BenchmarkReduce2_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce2(data, simpleAdd)
	}
}

// 方案3: 循环展开（每4个元素）
func Reduce3[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n == 1 {
		return ss[0]
	}

	result := ss[0]
	i := 1
	// 处理剩余元素，确保有4的倍数个元素可展开
	for i < n && ((n-i)%4) != 0 {
		result = f(result, ss[i])
		i++
	}

	// 4路展开
	for i+3 < n {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
		i += 4
	}

	// 处理剩余元素
	for i < n {
		result = f(result, ss[i])
		i++
	}

	return result
}

func BenchmarkReduce3_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce3(data, simpleAdd)
	}
}

func BenchmarkReduce3_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce3(data, simpleAdd)
	}
}

// 方案4: 使用指针传递（仅适用于指针类型）
func Reduce4[T any](ss []T, f func(*T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(&result, ss[i])
	}
	return result
}

func ptrAdd(a *int, b int) int {
	return *a + b
}

func BenchmarkReduce4_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce4(data, ptrAdd)
	}
}

func BenchmarkReduce4_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce4(data, ptrAdd)
	}
}

// 方案5: int类型特化版本
func ReduceInt(ss []int, f func(int, int) int) int {
	n := len(ss)
	if n == 0 {
		return 0
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

func BenchmarkReduceInt_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceInt(data, simpleAdd)
	}
}

func BenchmarkReduceInt_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceInt(data, simpleAdd)
	}
}

// 方案6: string类型特化版本
func ReduceString(ss []string, f func(string, string) string) string {
	n := len(ss)
	if n == 0 {
		return ""
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

func BenchmarkReduceString_Small(b *testing.B) {
	data := genStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceString(data, concatStr)
	}
}

func BenchmarkReduceString_Large(b *testing.B) {
	data := genStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReduceString(data, concatStr)
	}
}

// 方案7: 混合策略 - 小切片用简单循环，大切片用展开
func Reduce7[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}
	if n <= 4 {
		// 小切片：直接展开
		if n == 1 {
			return ss[0]
		}
		if n == 2 {
			return f(ss[0], ss[1])
		}
		if n == 3 {
			return f(f(ss[0], ss[1]), ss[2])
		}
		return f(f(f(ss[0], ss[1]), ss[2]), ss[3])
	}

	// 大切片：循环展开
	result := ss[0]
	i := 1
	for i+7 < n {
		result = f(result, ss[i])
		result = f(result, ss[i+1])
		result = f(result, ss[i+2])
		result = f(result, ss[i+3])
		result = f(result, ss[i+4])
		result = f(result, ss[i+5])
		result = f(result, ss[i+6])
		result = f(result, ss[i+7])
		i += 8
	}

	for i < n {
		result = f(result, ss[i])
		i++
	}

	return result
}

func BenchmarkReduce7_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

func BenchmarkReduce7_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

// 方案8: 针对复杂函数的优化（减少函数调用开销）
func Reduce8[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	// 对小切片直接返回
	if n <= 4 {
		switch n {
		case 1:
			return ss[0]
		case 2:
			return f(ss[0], ss[1])
		case 3:
			return f(f(ss[0], ss[1]), ss[2])
		case 4:
			return f(f(f(ss[0], ss[1]), ss[2]), ss[3])
		}
	}

	result := ss[0]
	for i := 1; i < n; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce8_Complex_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce8(data, simpleAdd)
	}
}

func BenchmarkReduce8_Complex_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce8(data, simpleAdd)
	}
}

// 方案9: 使用临时变量减少边界检查
func Reduce9[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	length := n
	for i := 1; i < length; i++ {
		result = f(result, ss[i])
	}
	return result
}

func BenchmarkReduce9_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce9(data, simpleAdd)
	}
}

func BenchmarkReduce9_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce9(data, simpleAdd)
	}
}

// 方案10: 无边界检查的循环（使用goto）
func Reduce10[T any](ss []T, f func(T, T) T) T {
	n := len(ss)
	if n == 0 {
		return *new(T)
	}

	result := ss[0]
	i := 1
	if i >= n {
		return result
	}

loop:
	result = f(result, ss[i])
	i++
	if i < n {
		goto loop
	}

	return result
}

func BenchmarkReduce10_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce10(data, simpleAdd)
	}
}

func BenchmarkReduce10_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce10(data, simpleAdd)
	}
}

// ==================== 测试复杂函数场景 ====================

func BenchmarkReduce_Current_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, simpleAdd)
	}
}

func BenchmarkReduce1_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, simpleAdd)
	}
}

func BenchmarkReduce7_Complex(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, simpleAdd)
	}
}

// ==================== 字符串测试 ====================

func BenchmarkReduce_Current_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce(data, concatStr)
	}
}

func BenchmarkReduce1_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce1(data, concatStr)
	}
}

func BenchmarkReduce7_String(b *testing.B) {
	data := genStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reduce7(data, concatStr)
	}
}

// ==================== Map 函数性能优化测试 ====================

// 基准方案：当前实现
func BenchmarkMap_Current_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap_Current_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap_Current_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

// 方案1：移除空检查分支
func BenchmarkMap1_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap1_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

func BenchmarkMap1_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		ret := make([]int, n)
		for j := 0; j < n; j++ {
			ret[j] = data[j] * 2
		}
		_ = ret
	}
}

// 方案2：循环展开2
func BenchmarkMap2_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap2_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap2_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 2 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
		}
		_ = ret
	}
}

// 方案3：循环展开4
func BenchmarkMap3_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap3_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

func BenchmarkMap3_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		for j := 0; j < n; j += 4 {
			ret[j] = data[j] * 2
			if j+1 < n {
				ret[j+1] = data[j+1] * 2
			}
			if j+2 < n {
				ret[j+2] = data[j+2] * 2
			}
			if j+3 < n {
				ret[j+3] = data[j+3] * 2
			}
		}
		_ = ret
	}
}

// 方案4：range循环对比
func BenchmarkMap4_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

func BenchmarkMap4_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

func BenchmarkMap4_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := make([]int, len(data))
		for idx, v := range data {
			ret[idx] = v * 2
		}
		_ = ret
	}
}

// 方案5：嵌套展开4
func BenchmarkMap5_Int_Small(b *testing.B) {
	data := genInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

func BenchmarkMap5_Int_Medium(b *testing.B) {
	data := genInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

func BenchmarkMap5_Int_Large(b *testing.B) {
	data := genInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := len(data)
		if n == 0 {
			continue
		}
		ret := make([]int, n)
		j := 0
		for j < n {
			if j+4 <= n {
				ret[j] = data[j] * 2
				ret[j+1] = data[j+1] * 2
				ret[j+2] = data[j+2] * 2
				ret[j+3] = data[j+3] * 2
				j += 4
			} else {
				ret[j] = data[j] * 2
				j++
			}
		}
		_ = ret
	}
}

// ==================== Reverse Benchmark Implementations ====================

// ReverseV1: 当前实现（小切片元素复制，大切片双指针）
func ReverseV1[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 32 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV2: 纯双指针交换（无分支）
func ReverseV2[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV3: 反向索引复制（无交换）
func ReverseV3[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = ss[n-1-i]
	}

	return result
}

// ReverseV4: 双指针 + 双向复制
func ReverseV4[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		result[i] = ss[j]
		result[j] = ss[i]
	}

	return result
}

// ReverseV5: 小切片阈值16
func ReverseV5[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 16 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV6: 小切片阈值64
func ReverseV6[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 64 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV7: 双指针 + 边界检查优化
func ReverseV7[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	left, right := 0, n-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

// ReverseV8: 预分配 + 索引反转
func ReverseV8[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)
	copy(result, ss)

	for i := n/2 - 1; i >= 0; i-- {
		opp := n - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}

// ReverseV9: 小切片阈值8
func ReverseV9[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 8 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ReverseV10: 小切片阈值128
func ReverseV10[T any](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	result := make([]T, n)

	if n < 128 {
		for i := 0; i < n; i++ {
			result[i] = ss[n-1-i]
		}
		return result
	}

	copy(result, ss)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ==================== Shuffle Benchmark Implementations ====================

// ShuffleV1: 当前实现（标准库 rand.Shuffle）
func ShuffleV1[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	stdrand.Shuffle(n, func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})

	return ss
}

// ShuffleV2: Fisher-Yates 手动实现
func ShuffleV2[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV3: Fisher-Yates 正向遍历
func ShuffleV3[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := 0; i < n-1; i++ {
		j := rand.IntN(n - i)
		ss[i], ss[i+j] = ss[i+j], ss[i]
	}

	return ss
}

// ShuffleV4: Fisher-Yates + 自交换检查
func ShuffleV4[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		if i != j {
			ss[i], ss[j] = ss[j], ss[i]
		}
	}

	return ss
}

// ShuffleV5: 局部变量优化
func ShuffleV5[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	s := ss
	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	return ss
}

// ShuffleV6: 小切片优化
func ShuffleV6[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	if n < 8 {
		for i := n - 1; i > 0; i-- {
			j := rand.IntN(i + 1)
			ss[i], ss[j] = ss[j], ss[i]
		}
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV7: 循环展开2个元素
func ShuffleV7[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV8: 分区洗牌
func ShuffleV8[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	mid := n / 2
	for i := mid - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}
	for i := n - 1; i > mid; i-- {
		j := rand.IntN(i - mid + 1)
		j += mid
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ShuffleV9: 块处理
func ShuffleV9[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i -= 2 {
		if i > 0 {
			j := rand.IntN(i + 1)
			ss[i], ss[j] = ss[j], ss[i]
			i--
			if i > 0 {
				j = rand.IntN(i + 1)
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}

	return ss
}

// ShuffleV10: 简化版本
func ShuffleV10[T any](ss []T) []T {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	for i := n - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

// ==================== Sort Benchmark Implementations ====================

// SortV1: 当前实现（小切片插入排序，大切片sort.Slice）
func SortV1[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV2: 纯 sort.Slice
func SortV2[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV3: 插入排序阈值16
func SortV3[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 16 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV4: 插入排序阈值32
func SortV4[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 32 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV5: 预排序检查
func SortV5[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if sorted[i] < sorted[i-1] {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV6: 预排序 + 插入排序
func SortV6[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if sorted[i] < sorted[i-1] {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV7: 二分插入排序
func SortV7[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			left, right := 0, i
			for left < right {
				mid := (left + right) / 2
				if sorted[mid] < key {
					left = mid + 1
				} else {
					right = mid
				}
			}
			for j := i; j > left; j-- {
				sorted[j] = sorted[j-1]
			}
			sorted[left] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV8: 插入排序阈值8
func SortV8[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV9: 插入排序阈值64
func SortV9[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 64 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// SortV10: 分层阈值（8/24/64）
func SortV10[T constraints.Ordered](ss []T) []T {
	n := len(ss)
	if n < 2 {
		return ss
	}

	sorted := make([]T, n)
	copy(sorted, ss)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	if n <= 64 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && sorted[j] > key {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// ==================== SortUsing Benchmark Implementations ====================

// SortUsingV1: 当前实现
func SortUsingV1[T any](slice []T, less func(T, T) bool) []T {
	if len(slice) < 2 {
		return slice
	}

	sorted := make([]T, len(slice))
	copy(sorted, slice)

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV2: 小切片插入排序
func SortUsingV2[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV3: 预排序检查
func SortUsingV3[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if less(sorted[i], sorted[i-1]) {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV4: 预排序 + 插入排序
func SortUsingV4[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	alreadySorted := true
	for i := 1; i < n; i++ {
		if less(sorted[i], sorted[i-1]) {
			alreadySorted = false
			break
		}
	}

	if alreadySorted {
		return sorted
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV5: 二分插入排序
func SortUsingV5[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			left, right := 0, i
			for left < right {
				mid := (left + right) / 2
				if less(sorted[mid], key) {
					left = mid + 1
				} else {
					right = mid
				}
			}
			for j := i; j > left; j-- {
				sorted[j] = sorted[j-1]
			}
			sorted[left] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV6: 插入排序阈值16
func SortUsingV6[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 16 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV7: 插入排序阈值32
func SortUsingV7[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 32 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV8: 插入排序阈值8
func SortUsingV8[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n <= 8 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// SortUsingV9: 减少闭包分配
func SortUsingV9[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	lessFunc := func(i, j int) bool {
		return less(sorted[i], sorted[j])
	}

	sort.Slice(sorted, lessFunc)

	return sorted
}

// SortUsingV10: 预检查 + 插入排序
func SortUsingV10[T any](slice []T, less func(T, T) bool) []T {
	n := len(slice)
	if n < 2 {
		return slice
	}

	sorted := make([]T, n)
	copy(sorted, slice)

	if n > 3 {
		if less(sorted[2], sorted[1]) || less(sorted[1], sorted[0]) {
			needsSort := false
			for i := 3; i < n; i++ {
				if less(sorted[i], sorted[i-1]) {
					needsSort = true
					break
				}
			}
			if !needsSort {
				return sorted
			}
		}
	}

	if n <= 24 {
		for i := 1; i < n; i++ {
			key := sorted[i]
			j := i - 1
			for j >= 0 && less(key, sorted[j]) {
				sorted[j+1] = sorted[j]
				j--
			}
			sorted[j+1] = key
		}
		return sorted
	}

	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

// ==================== Reverse Benchmarks ====================

func BenchmarkReverse_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]int{}, data...))
		}
	})
}

func BenchmarkReverse_String_Small(b *testing.B) {
	data := genStrings(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]string{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]string{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]string{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]string{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]string{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]string{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]string{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]string{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]string{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]string{}, data...))
		}
	})
}

func BenchmarkReverse_String_Large(b *testing.B) {
	data := genStrings(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV1(append([]string{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV2(append([]string{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV3(append([]string{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV4(append([]string{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV5(append([]string{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV6(append([]string{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV7(append([]string{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV8(append([]string{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV9(append([]string{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReverseV10(append([]string{}, data...))
		}
	})
}

// ==================== Shuffle Benchmarks ====================

func BenchmarkShuffle_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

func BenchmarkShuffle_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

func BenchmarkShuffle_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV1(d)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV2(d)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV3(d)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV4(d)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV5(d)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV6(d)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV7(d)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV8(d)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV9(d)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d := append([]int{}, data...)
			_ = ShuffleV10(d)
		}
	})
}

// ==================== Sort Benchmarks ====================

func BenchmarkSort_Int_Small(b *testing.B) {
	data := genInts(8)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Int_Medium(b *testing.B) {
	data := genInts(64)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Int_Large(b *testing.B) {
	data := genInts(1024)
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

func BenchmarkSort_Sorted_Large(b *testing.B) {
	data := make([]int, 1024)
	for i := range data {
		data[i] = i
	}
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV1(append([]int{}, data...))
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV2(append([]int{}, data...))
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV3(append([]int{}, data...))
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV4(append([]int{}, data...))
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV5(append([]int{}, data...))
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV6(append([]int{}, data...))
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV7(append([]int{}, data...))
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV8(append([]int{}, data...))
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV9(append([]int{}, data...))
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortV10(append([]int{}, data...))
		}
	})
}

// ==================== SortUsing Benchmarks ====================

func BenchmarkSortUsing_Int_Small(b *testing.B) {
	data := genInts(8)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}

func BenchmarkSortUsing_Int_Medium(b *testing.B) {
	data := genInts(64)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}

func BenchmarkSortUsing_Int_Large(b *testing.B) {
	data := genInts(1024)
	less := func(a, b int) bool { return a < b }
	b.Run("V1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV1(append([]int{}, data...), less)
		}
	})
	b.Run("V2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV2(append([]int{}, data...), less)
		}
	})
	b.Run("V3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV3(append([]int{}, data...), less)
		}
	})
	b.Run("V4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV4(append([]int{}, data...), less)
		}
	})
	b.Run("V5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV5(append([]int{}, data...), less)
		}
	})
	b.Run("V6", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV6(append([]int{}, data...), less)
		}
	})
	b.Run("V7", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV7(append([]int{}, data...), less)
		}
	})
	b.Run("V8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV8(append([]int{}, data...), less)
		}
	})
	b.Run("V9", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV9(append([]int{}, data...), less)
		}
	})
	b.Run("V10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SortUsingV10(append([]int{}, data...), less)
		}
	})
}

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
		s[i] = stdrand.Intn(100)
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
		p[i] = point{X: stdrand.Intn(100), Y: stdrand.Intn(100)}
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

// ==================== Join 基准测试 ====================

func BenchmarkJoin_Int_Small_Default_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Small_Comma_V1(b *testing.B) {
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, ",")
	}
}

func BenchmarkJoin_Int_Medium_Default_V1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Large_Default_V1(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_String_Small_Default_V1(b *testing.B) {
	data := make([]string, 10)
	for i := 0; i < 10; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_String_Medium_Dash_V1(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, "-")
	}
}

func BenchmarkJoin_String_Large_Space_V1(b *testing.B) {
	data := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = "test"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data, " ")
	}
}

func BenchmarkJoin_Int_Empty_V1(b *testing.B) {
	data := []int{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Float64_Large_Default_V1(b *testing.B) {
	data := make([]float64, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

func BenchmarkJoin_Int_Single_V1(b *testing.B) {
	data := []int{42}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(data)
	}
}

// ==================== 测试数据生成 ====================

// 生成整数切片（避免与 collection_ops_bench_test.go 冲突）
func genIntsEach(n int) []int {
	r := stdrand.New(stdrand.NewSource(42))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = r.Int()
	}
	return s
}

// ==================== 回调函数 ====================

// 简单回调（累加）
var sumResult int

func simpleSum(v int) {
	sumResult += v
}

// ==================== Each 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）- 使用 range
func Each1[T any](values []T, fn func(value T)) {
	for _, value := range values {
		fn(value)
	}
}

// 方案2: 使用索引循环
func Each2[T any](values []T, fn func(value T)) {
	n := len(values)
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案3: 预先检查空切片
func Each3[T any](values []T, fn func(value T)) {
	if len(values) == 0 {
		return
	}
	for _, value := range values {
		fn(value)
	}
}

// 方案4: 使用索引循环 + 空切片检查
func Each4[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案5: 循环展开（4路）
func Each5[T any](values []T, fn func(value T)) {
	n := len(values)
	i := 0

	// 处理不能被4整除的部分
	for i < n && (n-i)%4 != 0 {
		fn(values[i])
		i++
	}

	// 4路展开
	for i+3 < n {
		fn(values[i])
		fn(values[i+1])
		fn(values[i+2])
		fn(values[i+3])
		i += 4
	}

	// 处理剩余元素
	for i < n {
		fn(values[i])
		i++
	}
}

// 方案6: 循环展开（8路）
func Each6[T any](values []T, fn func(value T)) {
	n := len(values)
	i := 0

	// 处理不能被8整除的部分
	for i < n && (n-i)%8 != 0 {
		fn(values[i])
		i++
	}

	// 8路展开
	for i+7 < n {
		fn(values[i])
		fn(values[i+1])
		fn(values[i+2])
		fn(values[i+3])
		fn(values[i+4])
		fn(values[i+5])
		fn(values[i+6])
		fn(values[i+7])
		i += 8
	}

	// 处理剩余元素
	for i < n {
		fn(values[i])
		i++
	}
}

// 方案7: 小切片特殊处理
func Each7[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		if n >= 1 {
			fn(values[0])
		}
		if n >= 2 {
			fn(values[1])
		}
		if n >= 3 {
			fn(values[2])
		}
		if n >= 4 {
			fn(values[3])
		}
		return
	}

	// 大切片使用循环展开
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案8: 使用临时变量减少边界检查
func Each8[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	length := n
	for i := 0; i < length; i++ {
		fn(values[i])
	}
}

// 方案9: 混合策略 - 小切片用range，大切片用索引
func Each9[T any](values []T, fn func(value T)) {
	n := len(values)
	if n == 0 {
		return
	}

	// 小切片使用range（代码简洁）
	if n < 32 {
		for _, value := range values {
			fn(value)
		}
		return
	}

	// 大切片使用索引循环（性能更好）
	for i := 0; i < n; i++ {
		fn(values[i])
	}
}

// 方案10: 使用指针访问（适用于特定类型）
func Each10(values []int, fn func(*int)) {
	n := len(values)
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		fn(&values[i])
	}
}

func ptrSum(v *int) {
	sumResult += *v
}

// ==================== EachReverse 函数的 10 种优化方案 ====================

// 方案1: 当前实现（baseline）
func EachReverse1[T any](ss []T, f func(T)) {
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案2: 预先计算长度
func EachReverse2[T any](ss []T, f func(T)) {
	n := len(ss)
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案3: 空切片检查
func EachReverse3[T any](ss []T, f func(T)) {
	if len(ss) == 0 {
		return
	}
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案4: 预先计算长度 + 空切片检查
func EachReverse4[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案5: 循环展开（4路）
func EachReverse5[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1

	// 处理不能被4整除的部分
	for i >= 0 && ((n-1-i)%4) != 0 {
		f(ss[i])
		i--
	}

	// 4路展开
	for i-3 >= 0 {
		f(ss[i])
		f(ss[i-1])
		f(ss[i-2])
		f(ss[i-3])
		i -= 4
	}

	// 处理剩余元素
	for i >= 0 {
		f(ss[i])
		i--
	}
}

// 方案6: 循环展开（8路）
func EachReverse6[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1

	// 处理不能被8整除的部分
	for i >= 0 && ((n-1-i)%8) != 0 {
		f(ss[i])
		i--
	}

	// 8路展开
	for i-7 >= 0 {
		f(ss[i])
		f(ss[i-1])
		f(ss[i-2])
		f(ss[i-3])
		f(ss[i-4])
		f(ss[i-5])
		f(ss[i-6])
		f(ss[i-7])
		i -= 8
	}

	// 处理剩余元素
	for i >= 0 {
		f(ss[i])
		i--
	}
}

// 方案7: 小切片特殊处理
func EachReverse7[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		if n >= 1 {
			f(ss[n-1])
		}
		if n >= 2 {
			f(ss[n-2])
		}
		if n >= 3 {
			f(ss[n-3])
		}
		if n >= 4 {
			f(ss[n-4])
		}
		return
	}

	// 大切片使用常规循环
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案8: 使用临时变量减少边界检查
func EachReverse8[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	length := n
	for i := length - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案9: 混合策略 - 小和大切片不同处理
func EachReverse9[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	// 小切片直接展开
	if n <= 4 {
		for i := n - 1; i >= 0; i-- {
			f(ss[i])
		}
		return
	}

	// 大切片使用预先计算长度
	for i := n - 1; i >= 0; i-- {
		f(ss[i])
	}
}

// 方案10: 使用 while 循环风格（goto）
func EachReverse10[T any](ss []T, f func(T)) {
	n := len(ss)
	if n == 0 {
		return
	}

	i := n - 1
loop:
	f(ss[i])
	i--
	if i >= 0 {
		goto loop
	}
}

// ==================== 基准测试 ====================

// Each 测试 - 小切片
func BenchmarkEach1_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each1(data, simpleSum)
	}
}

func BenchmarkEach2_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each2(data, simpleSum)
	}
}

func BenchmarkEach3_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each3(data, simpleSum)
	}
}

func BenchmarkEach4_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each4(data, simpleSum)
	}
}

func BenchmarkEach5_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each5(data, simpleSum)
	}
}

func BenchmarkEach6_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each6(data, simpleSum)
	}
}

func BenchmarkEach7_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each7(data, simpleSum)
	}
}

func BenchmarkEach8_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each8(data, simpleSum)
	}
}

func BenchmarkEach9_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each9(data, simpleSum)
	}
}

func BenchmarkEach10_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each10(data, ptrSum)
	}
}

// Each 测试 - 大切片
func BenchmarkEach1_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each1(data, simpleSum)
	}
}

func BenchmarkEach2_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each2(data, simpleSum)
	}
}

func BenchmarkEach4_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each4(data, simpleSum)
	}
}

func BenchmarkEach5_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each5(data, simpleSum)
	}
}

func BenchmarkEach6_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each6(data, simpleSum)
	}
}

func BenchmarkEach9_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		Each9(data, simpleSum)
	}
}

// EachReverse 测试 - 小切片
func BenchmarkEachReverse1_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse1(data, simpleSum)
	}
}

func BenchmarkEachReverse2_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse2(data, simpleSum)
	}
}

func BenchmarkEachReverse4_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse4(data, simpleSum)
	}
}

func BenchmarkEachReverse9_Small(b *testing.B) {
	data := genIntsEach(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse9(data, simpleSum)
	}
}

// EachReverse 测试 - 大切片
func BenchmarkEachReverse1_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse1(data, simpleSum)
	}
}

func BenchmarkEachReverse2_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse2(data, simpleSum)
	}
}

func BenchmarkEachReverse4_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse4(data, simpleSum)
	}
}

func BenchmarkEachReverse9_Large(b *testing.B) {
	data := genIntsEach(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sumResult = 0
		EachReverse9(data, simpleSum)
	}
}

func BenchmarkFilter_Small_50Percent(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_50Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Large_50Percent(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n%2 == 0 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_10Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n < 100 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

func BenchmarkFilter_Medium_90Percent(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	predicate := func(n int) bool { return n < 900 }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Filter(data, predicate)
	}
}

// 生成测试数据
func genJoinInts(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = (i * 13) % 1000
	}
	return s
}

func genJoinStrings(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = strconv.Itoa((i * 13) % 1000)
	}
	return s
}

// ==================== Int 类型的基准测试 ====================

func BenchmarkJoin_Int_Small(b *testing.B) {
	data := genJoinInts(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Medium(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Large(b *testing.B) {
	data := genJoinInts(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

// ==================== String 类型的基准测试 ====================

func BenchmarkJoin_String_Small(b *testing.B) {
	data := genJoinStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_Medium(b *testing.B) {
	data := genJoinStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_Large(b *testing.B) {
	data := genJoinStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

// ==================== 不同分隔符的基准测试 ====================

func BenchmarkJoin_Int_Comma(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_Int_Dash(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "-")
	}
}

func BenchmarkJoin_Int_Pipe(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "|")
	}
}

// ==================== 空分隔符测试 ====================

func BenchmarkJoin_Int_EmptyGlue(b *testing.B) {
	data := genJoinInts(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "")
	}
}

func BenchmarkJoin_String_EmptyGlue(b *testing.B) {
	data := genJoinStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, "")
	}
}

// ==================== 空切片测试 ====================

func BenchmarkJoin_Int_EmptySlice(b *testing.B) {
	data := genJoinInts(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}

func BenchmarkJoin_String_EmptySlice(b *testing.B) {
	data := genJoinStrings(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Join(data, ",")
	}
}
