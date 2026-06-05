package candy

import (
	"testing"
)

// Benchmark Chunk 函数的各种实现方案

// 方案1: 当前实现 (混合策略)
func BenchmarkChunk_Current_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(data, 10)
	}
}

// 方案2: 纯 range 循环
func ChunkRange[T any](ss []T, size int) [][]T {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}
	if size >= len(ss) {
		return [][]T{ss}
	}

	ret := make([][]T, 0, (len(ss)+size-1)/size)
	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}
		ret = append(ret, ss[i:end:end])
	}
	return ret
}

func BenchmarkChunk_Range_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkRange(data, 10)
	}
}

// 方案3: 传统 for 循环
func ChunkFor[T any](ss []T, size int) [][]T {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}
	if size >= len(ss) {
		return [][]T{ss}
	}

	n := len(ss)
	ret := make([][]T, 0, (n+size-1)/size)
	i := 0
	for i < n {
		end := i + size
		if end > n {
			end = n
		}
		ret = append(ret, ss[i:end:end])
		i = end
	}
	return ret
}

func BenchmarkChunk_For_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFor(data, 10)
	}
}

// 方案4: 完全切片优化（使用三参数切片）
func ChunkFullSlice[T any](ss []T, size int) [][]T {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}
	if size >= len(ss) {
		return [][]T{ss}
	}

	n := len(ss)
	chunkCount := (n + size - 1) / size
	ret := make([][]T, chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * size
		end := start + size
		if end > n {
			end = n
		}
		ret[i] = ss[start:end:end]
	}
	return ret
}

func BenchmarkChunk_FullSlice_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFullSlice(data, 10)
	}
}

// 方案5: 预分配后直接 append（不使用完全切片）
func ChunkAppend[T any](ss []T, size int) [][]T {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}
	if size >= len(ss) {
		return [][]T{ss}
	}

	ret := make([][]T, 0, (len(ss)+size-1)/size)
	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}
		ret = append(ret, ss[i:end])
		i = end
	}
	return ret
}

func BenchmarkChunk_Append_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkAppend(data, 10)
	}
}

// 中等数据集测试
func BenchmarkChunk_Current_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(data, 10)
	}
}

func BenchmarkChunk_Range_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkRange(data, 10)
	}
}

func BenchmarkChunk_For_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFor(data, 10)
	}
}

func BenchmarkChunk_FullSlice_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFullSlice(data, 10)
	}
}

func BenchmarkChunk_Append_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkAppend(data, 10)
	}
}

// 大数据集测试
func BenchmarkChunk_Current_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(data, 100)
	}
}

func BenchmarkChunk_Range_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkRange(data, 100)
	}
}

func BenchmarkChunk_For_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFor(data, 100)
	}
}

func BenchmarkChunk_FullSlice_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFullSlice(data, 100)
	}
}

func BenchmarkChunk_Append_Large(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkAppend(data, 100)
	}
}

// 特殊情况：size=1
func BenchmarkChunk_Current_Size1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(data, 1)
	}
}

func BenchmarkChunk_Range_Size1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkRange(data, 1)
	}
}

func BenchmarkChunk_For_Size1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFor(data, 1)
	}
}

func BenchmarkChunk_FullSlice_Size1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFullSlice(data, 1)
	}
}

func BenchmarkChunk_Append_Size1(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkAppend(data, 1)
	}
}

// 特殊情况：size >= len(ss)
func BenchmarkChunk_Current_SizeLarge(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Chunk(data, 200)
	}
}

func BenchmarkChunk_Range_SizeLarge(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkRange(data, 200)
	}
}

func BenchmarkChunk_For_SizeLarge(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFor(data, 200)
	}
}

func BenchmarkChunk_FullSlice_SizeLarge(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkFullSlice(data, 200)
	}
}

func BenchmarkChunk_Append_SizeLarge(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ChunkAppend(data, 200)
	}
}
