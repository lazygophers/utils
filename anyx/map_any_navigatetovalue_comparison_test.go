package anyx

import (
	"fmt"
	"testing"
)

// navigateToValue 性能优化对比测试
// 对比当前实现和不同优化方案的性能

// ============================================================
// 场景 1: 简单 map 键访问（最常见场景）
// ============================================================

func BenchmarkCompare_SimpleMapKey_Current(b *testing.B) {
	data := map[string]any{"name": "value"}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_SimpleMapKey_Optimized(b *testing.B) {
	data := map[string]any{"name": "value"}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_SimpleMapKey_OptimizedV2(b *testing.B) {
	data := map[string]any{"name": "value"}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV2(data, part)
	}
}

func BenchmarkCompare_SimpleMapKey_OptimizedV3(b *testing.B) {
	data := map[string]any{"name": "value"}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV3(data, part)
	}
}

func BenchmarkCompare_SimpleMapKey_OptimizedV4(b *testing.B) {
	data := map[string]any{"name": "value"}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 2: []any 数组索引访问
// ============================================================

func BenchmarkCompare_ArrayIndexAny_Current(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_ArrayIndexAny_Optimized(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_ArrayIndexAny_OptimizedV2(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV2(data, part)
	}
}

func BenchmarkCompare_ArrayIndexAny_OptimizedV3(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV3(data, part)
	}
}

func BenchmarkCompare_ArrayIndexAny_OptimizedV4(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 3: []string 数组索引访问
// ============================================================

func BenchmarkCompare_ArrayIndexString_Current(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_ArrayIndexString_Optimized(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_ArrayIndexString_OptimizedV4(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 4: []int 数组索引访问
// ============================================================

func BenchmarkCompare_ArrayIndexInt_Current(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_ArrayIndexInt_Optimized(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_ArrayIndexInt_OptimizedV4(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 5: 大索引值
// ============================================================

func BenchmarkCompare_LargeIndex_Current(b *testing.B) {
	data := make([]any, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	part := "[999]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_LargeIndex_Optimized(b *testing.B) {
	data := make([]any, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	part := "[999]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_LargeIndex_OptimizedV4(b *testing.B) {
	data := make([]any, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	part := "[999]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 6: 大型 map 访问
// ============================================================

func BenchmarkCompare_LargeMap_Current(b *testing.B) {
	data := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		data[fmt.Sprintf("key%d", i)] = i
	}
	part := "key500"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_LargeMap_Optimized(b *testing.B) {
	data := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		data[fmt.Sprintf("key%d", i)] = i
	}
	part := "key500"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_LargeMap_OptimizedV4(b *testing.B) {
	data := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		data[fmt.Sprintf("key%d", i)] = i
	}
	part := "key500"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 7: 键不存在（错误路径）
// ============================================================

func BenchmarkCompare_KeyNotFound_Current(b *testing.B) {
	data := map[string]any{"existing": "value"}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_KeyNotFound_Optimized(b *testing.B) {
	data := map[string]any{"existing": "value"}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_KeyNotFound_OptimizedV4(b *testing.B) {
	data := map[string]any{"existing": "value"}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 8: 索引越界（错误路径）
// ============================================================

func BenchmarkCompare_IndexOutOfRange_Current(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_IndexOutOfRange_Optimized(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_IndexOutOfRange_OptimizedV4(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 9: 无效索引格式（错误路径）
// ============================================================

func BenchmarkCompare_InvalidIndex_Current(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_InvalidIndex_Optimized(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_InvalidIndex_OptimizedV4(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}

// ============================================================
// 场景 10: map[any]any 访问
// ============================================================

func BenchmarkCompare_MapAnyAny_Current(b *testing.B) {
	data := map[any]any{"key": "value"}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkCompare_MapAnyAny_Optimized(b *testing.B) {
	data := map[any]any{"key": "value"}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

func BenchmarkCompare_MapAnyAny_OptimizedV4(b *testing.B) {
	data := map[any]any{"key": "value"}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimizedV4(data, part)
	}
}
