package anyx

import (
	"fmt"
	"testing"
)

// 简化的基准测试，避免依赖问题

// BenchmarkNavigateToValue_Current_ 系列测试当前实现

func BenchmarkNavigateToValue_Current_SimpleMapKey(b *testing.B) {
	data := map[string]any{
		"name": "value",
	}
	part := "name"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexAny(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexString(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexInt(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexInt64(b *testing.B) {
	data := []int64{1, 2, 3, 4, 5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexFloat64(b *testing.B) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexBool(b *testing.B) {
	data := []bool{true, false, true, false, true}
	part := "[2]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_ArrayIndexMap(b *testing.B) {
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

func BenchmarkNavigateToValue_Current_LargeIndex(b *testing.B) {
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

func BenchmarkNavigateToValue_Current_EmptyKey(b *testing.B) {
	data := map[string]any{
		"": "empty-value",
	}
	part := ""

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_MapAnyAny(b *testing.B) {
	data := map[any]any{
		"key": "value",
	}
	part := "key"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_KeyNotFound(b *testing.B) {
	data := map[string]any{
		"existing": "value",
	}
	part := "nonexistent"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_IndexOutOfRange(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_InvalidIndex(b *testing.B) {
	data := []any{"a", "b", "c"}
	part := "[abc]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = navigateToValue(data, part)
	}
}

func BenchmarkNavigateToValue_Current_LargeMap(b *testing.B) {
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

// ============================================================
// 测试 parseIndex 的性能
// ============================================================

func BenchmarkParseIndex_Valid(b *testing.B) {
	indexStr := "123"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(indexStr)
	}
}

func BenchmarkParseIndex_LargeNumber(b *testing.B) {
	indexStr := "999999"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(indexStr)
	}
}

func BenchmarkParseIndex_Negative(b *testing.B) {
	indexStr := "-456"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(indexStr)
	}
}

func BenchmarkParseIndex_SingleDigit(b *testing.B) {
	indexStr := "5"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(indexStr)
	}
}

// ============================================================
// 测试 accessArrayIndex 的性能
// ============================================================

func BenchmarkAccessArrayIndex_AnySlice(b *testing.B) {
	data := []any{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_StringSlice(b *testing.B) {
	data := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_IntSlice(b *testing.B) {
	data := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_Int64Slice(b *testing.B) {
	data := []int64{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_Float64Slice(b *testing.B) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_BoolSlice(b *testing.B) {
	data := []bool{true, false, true, false, true}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "2")
	}
}

func BenchmarkAccessArrayIndex_MapSlice(b *testing.B) {
	data := []map[string]any{
		{"key": "value1"},
		{"key": "value2"},
		{"key": "value3"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessArrayIndex(data, "1")
	}
}

// ============================================================
// 测试 accessMapKey 的性能
// ============================================================

func BenchmarkAccessMapKey_MapStringAny(b *testing.B) {
	data := map[string]any{
		"key": "value",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKey(data, "key")
	}
}

func BenchmarkAccessMapKey_MapAnyAny(b *testing.B) {
	data := map[any]any{
		"key": "value",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKey(data, "key")
	}
}

func BenchmarkAccessMapKey_KeyNotFound(b *testing.B) {
	data := map[string]any{
		"existing": "value",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = accessMapKey(data, "nonexistent")
	}
}
