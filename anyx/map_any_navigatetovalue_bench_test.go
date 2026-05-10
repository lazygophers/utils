package anyx

import (
	"testing"
)

// Benchmark NavigateToValue - 不同场景的全面性能测试

// 场景 1: 简单 map 键访问
func BenchmarkNavigateToValue_SimpleMapKey(b *testing.B) {
	data := map[string]any{
		"name": "value",
	}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 2: 数组索引访问 - []any
func BenchmarkNavigateToValue_ArrayIndexAny(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 3: 数组索引访问 - []string
func BenchmarkNavigateToValue_ArrayIndexString(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 4: 数组索引访问 - []int
func BenchmarkNavigateToValue_ArrayIndexInt(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 5: 数组索引访问 - []int64
func BenchmarkNavigateToValue_ArrayIndexInt64(b *testing.B) {
	data := []int64{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 6: 数组索引访问 - []float64
func BenchmarkNavigateToValue_ArrayIndexFloat64(b *testing.B) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 7: 数组索引访问 - []bool
func BenchmarkNavigateToValue_ArrayIndexBool(b *testing.B) {
	data := []bool{true, false, true, false, true}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 8: 数组索引访问 - []map[string]any
func BenchmarkNavigateToValue_ArrayIndexMap(b *testing.B) {
	data := []map[string]any{
		{"key": "value1"},
		{"key": "value2"},
		{"key": "value3"},
	}
	part := "[1]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 9: 负数索引
func BenchmarkNavigateToValue_NegativeIndex(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[-2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 10: 大索引值
func BenchmarkNavigateToValue_LargeIndex(b *testing.B) {
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

// 场景 11: 空键访问
func BenchmarkNavigateToValue_EmptyKey(b *testing.B) {
	data := map[string]any{
		"": "empty-value",
	}
	part := ""

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 12: 空 part（默认行为）
func BenchmarkNavigateToValue_EmptyPart(b *testing.B) {
	data := map[string]any{
		"": "value",
	}
	part := ""

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 13: map[any]any 访问
func BenchmarkNavigateToValue_MapAnyAny(b *testing.B) {
	data := map[any]any{
		"key": "value",
	}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 14: 无效类型访问（错误路径）
func BenchmarkNavigateToValue_InvalidType(b *testing.B) {
	data := "not-a-map-or-slice"
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 15: 索引越界（错误路径）
func BenchmarkNavigateToValue_IndexOutOfRange(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 16: 无效索引格式（错误路径）
func BenchmarkNavigateToValue_InvalidIndexFormat(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 17: 键不存在（错误路径）
func BenchmarkNavigateToValue_KeyNotFound(b *testing.B) {
	data := map[string]any{
		"existing": "value",
	}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 18: 边界检查 - 索引 0
func BenchmarkNavigateToValue_FirstIndex(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[0]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 19: 边界检查 - 最后索引
func BenchmarkNavigateToValue_LastIndex(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[4]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 20: 混合场景 - 先 map 后 array
func BenchmarkNavigateToValue_MapThenArray(b *testing.B) {
	data := map[string]any{
		"items": []any{"a", "b", "c"},
	}
	part := "items"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 21: 大型 map 访问
func BenchmarkNavigateToValue_LargeMap(b *testing.B) {
	data := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		data[string(rune(i))] = i
	}
	part := "x" // ASCII 120

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 22: 复杂嵌套结构
func BenchmarkNavigateToValue_NestedStructure(b *testing.B) {
	data := map[string]any{
		"level1": map[string]any{
			"level2": "value",
		},
	}
	part := "level1"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 23: 多字节索引
func BenchmarkNavigateToValue_MultiDigitIndex(b *testing.B) {
	data := make([]any, 1000)
	part := "[999]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 24: 单字符键
func BenchmarkNavigateToValue_SingleCharKey(b *testing.B) {
	data := map[string]any{
		"a": "value",
	}
	part := "a"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 25: 长键名
func BenchmarkNavigateToValue_LongKey(b *testing.B) {
	data := map[string]any{
		"this-is-a-very-long-key-name-for-testing-purposes": "value",
	}
	part := "this-is-a-very-long-key-name-for-testing-purposes"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 26: 空数组访问
func BenchmarkNavigateToValue_EmptySlice(b *testing.B) {
	data := []any{}
	part := "[0]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 27: 空映射访问
func BenchmarkNavigateToValue_EmptyMap(b *testing.B) {
	data := map[string]any{}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 28: nil 数据访问（错误路径）
func BenchmarkNavigateToValue_NilData(b *testing.B) {
	var data []any = nil
	part := "[0]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 29: 带方括号但不是索引的键
func BenchmarkNavigateToValue_KeyWithBrackets(b *testing.B) {
	data := map[string]any{
		"[key]": "value",
	}
	part := "[key]" // 这会被当作索引处理，会失败

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// 场景 30: 并发安全的 map 访问
func BenchmarkNavigateToValue_ConcurrentMapAccess(b *testing.B) {
	data := map[string]any{
		"key": "value",
	}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

// ============================================================
// 优化版本的基准测试（待实现）- 暂时注释
// ============================================================

/*
// BenchmarkNavigateToValueOptimized_SimpleMapKey 测试优化版本的简单键访问
func BenchmarkNavigateToValueOptimized_SimpleMapKey(b *testing.B) {
	data := map[string]any{
		"name": "value",
	}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_ArrayIndex 测试优化版本的数组索引
func BenchmarkNavigateToValueOptimized_ArrayIndex(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_StringSlice 测试优化版本的字符串切片
func BenchmarkNavigateToValueOptimized_StringSlice(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_NegativeIndex 测试优化版本的负数索引
func BenchmarkNavigateToValueOptimized_NegativeIndex(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[-2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_LargeIndex 测试优化版本的大索引
func BenchmarkNavigateToValueOptimized_LargeIndex(b *testing.B) {
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

// BenchmarkNavigateToValueOptimized_MapAnyAny 测试优化版本的 map[any]any
func BenchmarkNavigateToValueOptimized_MapAnyAny(b *testing.B) {
	data := map[any]any{
		"key": "value",
	}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_KeyNotFound 测试优化版本的键不存在
func BenchmarkNavigateToValueOptimized_KeyNotFound(b *testing.B) {
	data := map[string]any{
		"existing": "value",
	}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_IndexOutOfRange 测试优化版本的索引越界
func BenchmarkNavigateToValueOptimized_IndexOutOfRange(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_InvalidIndex 测试优化版本的无效索引
func BenchmarkNavigateToValueOptimized_InvalidIndex(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_EmptyKey 测试优化版本的空键
func BenchmarkNavigateToValueOptimized_EmptyKey(b *testing.B) {
	data := map[string]any{
		"": "empty-value",
	}
	part := ""

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}

// BenchmarkNavigateToValueOptimized_LargeMap 测试优化版本的大型 map
func BenchmarkNavigateToValueOptimized_LargeMap(b *testing.B) {
	data := make(map[string]any, 1000)
	for i := 0; i < 1000; i++ {
		data[string(rune(i))] = i
	}
	part := "x"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValueOptimized(data, part)
	}
}
*/
