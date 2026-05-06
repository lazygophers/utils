package candy

import (
	"fmt"
	"testing"
)

// 性能验证测试 - 确保优化后的代码能正确处理大量数据
func TestPluckPerformance_Int32(t *testing.T) {
	data := generateBenchData(10000)

	// 测试能正确处理大量数据
	iterations := 100
	for i := 0; i < iterations; i++ {
		result := PluckInt32(data, "Age")
		if len(result) != 10000 {
			t.Fatalf("expected 10000 results, got %d", len(result))
		}
	}
}

func TestPluckPerformance_Int64(t *testing.T) {
	data := generateBenchData(10000)

	iterations := 100
	for i := 0; i < iterations; i++ {
		result := PluckInt64(data, "Score")
		if len(result) != 10000 {
			t.Fatalf("expected 10000 results, got %d", len(result))
		}
	}
}

func TestPluckPerformance_Uint32(t *testing.T) {
	data := generateBenchData(10000)

	iterations := 100
	for i := 0; i < iterations; i++ {
		result := PluckUint32(data, "Count")
		if len(result) != 10000 {
			t.Fatalf("expected 10000 results, got %d", len(result))
		}
	}
}

func TestPluckPerformance_Uint64(t *testing.T) {
	data := generateBenchData(10000)

	iterations := 100
	for i := 0; i < iterations; i++ {
		result := PluckUint64(data, "Total")
		if len(result) != 10000 {
			t.Fatalf("expected 10000 results, got %d", len(result))
		}
	}
}

func TestPluckPerformance_StringSlice(t *testing.T) {
	data := generateBenchData(10000)

	iterations := 100
	for i := 0; i < iterations; i++ {
		result := PluckStringSlice(data, "Tags")
		if len(result) != 10000 {
			t.Fatalf("expected 10000 results, got %d", len(result))
		}
	}
}

// 测试缓存功能是否正常工作
func TestPluckCache_Int32(t *testing.T) {
	data := generateBenchData(100)

	// 第一次调用会缓存
	result1 := PluckInt32(data, "Age")

	// 第二次调用应该使用缓存
	result2 := PluckInt32(data, "Age")

	if len(result1) != len(result2) {
		t.Fatalf("cache test failed: results have different lengths")
	}

	for i := range result1 {
		if result1[i] != result2[i] {
			t.Fatalf("cache test failed: results differ at index %d", i)
		}
	}
}

// 测试不同类型的数据结构
func TestPluckCache_MultipleTypes(t *testing.T) {
	type Type1 struct {
		Value int32
	}
	type Type2 struct {
		Value int32
	}

	data1 := []Type1{{Value: 1}, {Value: 2}}
	data2 := []Type2{{Value: 3}, {Value: 4}}

	result1 := PluckInt32(data1, "Value")
	result2 := PluckInt32(data2, "Value")

	if len(result1) != 2 || result1[0] != 1 || result1[1] != 2 {
		t.Fatalf("Type1 failed: got %v", result1)
	}
	if len(result2) != 2 || result2[0] != 3 || result2[1] != 4 {
		t.Fatalf("Type2 failed: got %v", result2)
	}
}

// 简单的性能对比示例
func ExamplePluckInt32() {
	data := generateBenchData(1000)
	result := PluckInt32(data, "Age")
	fmt.Printf("Extracted %d int32 values\n", len(result))
	// Output: Extracted 1000 int32 values
}
