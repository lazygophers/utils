package candy

import (
	"math/rand"
	"testing"
)

// ==================== 测试数据生成 ====================

// 生成整数切片
func genIntsPred(n int) []int {
	r := rand.New(rand.NewSource(42))
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
