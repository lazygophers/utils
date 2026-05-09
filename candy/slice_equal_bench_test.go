package candy

import "testing"

// Benchmark SliceEqual 函数的各种实现方案

// 方案1: 当前实现（小切片用双重循环，大切片用 map）
func BenchmarkSliceEqual_Current_Small(b *testing.B) {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	c := make([]int, 10)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqual(a, c)
	}
}

// 方案2: 始终使用 map（无小切片优化）
func SliceEqualAlwaysMap[T any](a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	am := make(map[any]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		vAny := any(v)
		if count, ok := am[vAny]; !ok || count == 0 {
			return false
		}
		am[vAny]--
	}

	return true
}

func BenchmarkSliceEqual_AlwaysMap_Small(b *testing.B) {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	c := make([]int, 10)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualAlwaysMap(a, c)
	}
}

// 方案3: 优化的双重循环（使用索引）
func SliceEqualDoubleLoop[T any](a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	// 使用标记数组避免 map 开销
	matched := make([]bool, len(b))
	for i := 0; i < len(a); i++ {
		va := a[i]
		found := false
		for j := 0; j < len(b); j++ {
			if !matched[j] {
				vaAny := any(va)
				vbAny := any(b[j])
				if vaAny == vbAny {
					matched[j] = true
					found = true
					break
				}
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func BenchmarkSliceEqual_DoubleLoop_Small(b *testing.B) {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	c := make([]int, 10)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualDoubleLoop(a, c)
	}
}

// 方案4: 索引循环优化（当前方案的改进版）
func SliceEqualOptimized[T any](a, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	// 小切片优化：使用双重循环避免 map 开销
	const smallSliceThreshold = 32
	if len(a) < smallSliceThreshold {
		matched := make([]bool, len(b))
		for i := 0; i < len(a); i++ {
			va := a[i]
			found := false
			for j := 0; j < len(b); j++ {
				if !matched[j] {
					vaAny := any(va)
					vbAny := any(b[j])
					if vaAny == vbAny {
						matched[j] = true
						found = true
						break
					}
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	// 大切片：使用 map 计数
	am := make(map[any]int, len(a))
	for i := 0; i < len(a); i++ {
		am[a[i]]++
	}

	for i := 0; i < len(b); i++ {
		vAny := any(b[i])
		if count, ok := am[vAny]; !ok || count == 0 {
			return false
		}
		am[vAny]--
	}

	return true
}

func BenchmarkSliceEqual_Optimized_Small(b *testing.B) {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	c := make([]int, 10)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOptimized(a, c)
	}
}

// 中等数据集测试（小于32）
func BenchmarkSliceEqual_Current_MediumSmall(b *testing.B) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}
	c := make([]int, 20)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqual(a, c)
	}
}

func BenchmarkSliceEqual_AlwaysMap_MediumSmall(b *testing.B) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}
	c := make([]int, 20)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualAlwaysMap(a, c)
	}
}

func BenchmarkSliceEqual_Optimized_MediumSmall(b *testing.B) {
	a := make([]int, 20)
	for i := 0; i < 20; i++ {
		a[i] = i
	}
	c := make([]int, 20)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOptimized(a, c)
	}
}

// 大数据集测试（大于32）
func BenchmarkSliceEqual_Current_Large(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqual(a, c)
	}
}

func BenchmarkSliceEqual_AlwaysMap_Large(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualAlwaysMap(a, c)
	}
}

func BenchmarkSliceEqual_Optimized_Large(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOptimized(a, c)
	}
}

// 超大数据集测试
func BenchmarkSliceEqual_Current_VeryLarge(b *testing.B) {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	c := make([]int, 1000)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqual(a, c)
	}
}

func BenchmarkSliceEqual_AlwaysMap_VeryLarge(b *testing.B) {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	c := make([]int, 1000)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualAlwaysMap(a, c)
	}
}

func BenchmarkSliceEqual_Optimized_VeryLarge(b *testing.B) {
	a := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		a[i] = i
	}
	c := make([]int, 1000)
	copy(c, a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOptimized(a, c)
	}
}

// 不相等情况测试
func BenchmarkSliceEqual_Current_NotEqual(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	for i := 0; i < 100; i++ {
		c[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqual(a, c)
	}
}

func BenchmarkSliceEqual_AlwaysMap_NotEqual(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	for i := 0; i < 100; i++ {
		c[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualAlwaysMap(a, c)
	}
}

func BenchmarkSliceEqual_Optimized_NotEqual(b *testing.B) {
	a := make([]int, 100)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	c := make([]int, 100)
	for i := 0; i < 100; i++ {
		c[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOptimized(a, c)
	}
}
