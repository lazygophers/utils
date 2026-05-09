package candy

import "testing"

// Benchmark Top 函数的各种实现方案

// 方案1: 当前实现（使用 copy）
func BenchmarkTop_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(data, 10)
	}
}

// 方案2: 使用切片复制（不使用 copy 函数）
func TopSlice[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret := make([]T, n)
	for i := 0; i < n; i++ {
		ret[i] = ss[i]
	}
	return ret
}

func BenchmarkTop_Slice_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopSlice(data, 10)
	}
}

// 方案3: 直接返回切片（引用）
func TopRef[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	return ss[:n]
}

func BenchmarkTop_Ref_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopRef(data, 10)
	}
}

// 方案4: 使用 append（优化预分配）
func TopAppend[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret := make([]T, 0, n)
	for i := 0; i < n; i++ {
		ret = append(ret, ss[i])
	}
	return ret
}

func BenchmarkTop_Append_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopAppend(data, 10)
	}
}

// 方案5: 完全切片优化
func TopFullSlice[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret := make([]T, n)
	copy(ret, ss[:n:n])
	return ret
}

func BenchmarkTop_FullSlice_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopFullSlice(data, 10)
	}
}

// 中等数据集测试
func BenchmarkTop_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(data, 100)
	}
}

func BenchmarkTop_Slice_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopSlice(data, 100)
	}
}

func BenchmarkTop_Append_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopAppend(data, 100)
	}
}

// 大数据集测试
func BenchmarkTop_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Top(data, 1000)
	}
}

func BenchmarkTop_Slice_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopSlice(data, 1000)
	}
}

func BenchmarkTop_Append_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopAppend(data, 1000)
	}
}

// Benchmark Bottom 函数的各种实现方案

// 方案1: 当前实现
func BenchmarkBottom_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 10)
	}
}

// 方案2: 使用切片复制
func BottomSlice[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret := make([]T, n)
	start := l - n
	for i := 0; i < n; i++ {
		ret[i] = ss[start+i]
	}
	return ret
}

func BenchmarkBottom_Slice_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomSlice(data, 10)
	}
}

// 方案3: 使用 append
func BottomAppend[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}

	ret := make([]T, 0, n)
	start := l - n
	for i := start; i < l; i++ {
		ret = append(ret, ss[i])
	}
	return ret
}

func BenchmarkBottom_Append_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomAppend(data, 10)
	}
}

// 方案4: 直接返回切片引用
func BottomRef[T any](ss []T, n int) []T {
	l := len(ss)
	if n <= 0 {
		return []T{}
	}
	if n > l {
		n = l
	}
	return ss[l-n:]
}

func BenchmarkBottom_Ref_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomRef(data, 10)
	}
}

// 中等数据集测试
func BenchmarkBottom_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 100)
	}
}

func BenchmarkBottom_Slice_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomSlice(data, 100)
	}
}

func BenchmarkBottom_Append_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomAppend(data, 100)
	}
}

// 大数据集测试
func BenchmarkBottom_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bottom(data, 1000)
	}
}

func BenchmarkBottom_Slice_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomSlice(data, 1000)
	}
}

func BenchmarkBottom_Append_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomAppend(data, 1000)
	}
}
