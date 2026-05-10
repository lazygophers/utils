package anyx

import (
	"strconv"
	"testing"

	"github.com/lazygophers/utils/candy"
)

// 基准测试数据
var benchData = map[string]interface{}{
	"string_slice":    []string{"a", "b", "c", "d", "e"},
	"int_slice":       []int{1, 2, 3, 4, 5},
	"int64_slice":     []int64{1, 2, 3, 4, 5},
	"uint64_slice":    []uint64{1, 2, 3, 4, 5},
	"float64_slice":   []float64{1.1, 2.2, 3.3, 4.4, 5.5},
	"bool_slice":      []bool{true, false, true, false, true},
	"interface_slice": []interface{}{1, "a", true, 2, "b"},
}

func setupBenchMap() *MapAny {
	return NewMap(benchData)
}

// ==================== 原始实现 ====================
func GetStringSliceOrig(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch val.(type) {
	case []bool, []int, []int8, []int16, []int32, []int64,
		[]uint, []uint8, []uint16, []uint32, []uint64,
		[]float32, []float64, []string, [][]byte, []interface{}:
		return candy.ToStringSlice(val)
	default:
		return []string{}
	}
}

// ==================== 方案 1: 内联类型断言 ====================
func GetStringSliceV1(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []int:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []float64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = candy.ToString(x[i])
		}
		return result
	case []bool:
		result := make([]string, len(x))
		for i := range x {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []interface{}:
		result := make([]string, len(x))
		for i := range x {
			result[i] = candy.ToString(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 2: 移除 nil 检查 ====================
func GetStringSliceV2(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []int:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []float64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatFloat(x[i], 'f', -1, 64)
		}
		return result
	case []bool:
		result := make([]string, len(x))
		for i := range x {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []interface{}:
		result := make([]string, len(x))
		for i := range x {
			result[i] = candy.ToString(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 3: 预分配常量字符串 ====================
var (
	boolTrue  = "1"
	boolFalse = "0"
)

func GetStringSliceV3(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []int64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []bool:
		result := make([]string, len(x))
		for i := range x {
			if x[i] {
				result[i] = boolTrue
			} else {
				result[i] = boolFalse
			}
		}
		return result
	case []interface{}:
		result := make([]string, len(x))
		for i := range x {
			result[i] = candy.ToString(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 4: 优化类型断言顺序 ====================
func GetStringSliceV4(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string: // 最常见
		return x
	case []interface{}: // 第二常见
		result := make([]string, len(x))
		for i := range x {
			result[i] = candy.ToString(x[i])
		}
		return result
	case []int64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		result := make([]string, len(x))
		for i := range x {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []bool:
		result := make([]string, len(x))
		for i := range x {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 5: 使用索引循环 ====================
func GetStringSliceV5(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []int64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []bool:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []interface{}:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 6: 快速路径分离 ====================
func GetStringSliceV6(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	if x, ok := val.([]string); ok {
		return x
	}
	switch x := val.(type) {
	case []int64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []bool:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []interface{}:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 7: 完整展开 ====================
func GetStringSliceV7(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []int:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int8:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int16:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int32:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatInt(int64(x[i]), 10)
		}
		return result
	case []int64:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint8:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint16:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint32:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatUint(uint64(x[i]), 10)
		}
		return result
	case []uint64:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []float32:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatFloat(float64(x[i]), 'f', -1, 64)
		}
		return result
	case []float64:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = strconv.FormatFloat(x[i], 'f', -1, 64)
		}
		return result
	case []bool:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			if x[i] {
				result[i] = "1"
			} else {
				result[i] = "0"
			}
		}
		return result
	case []interface{}:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	case [][]byte:
		result := make([]string, len(x))
		for i := 0; i < len(x); i++ {
			result[i] = string(x[i])
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 8: 综合优化 ====================
func GetStringSliceV8(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []interface{}:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	case []int64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	case []bool:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			if x[i] {
				result[i] = boolTrue
			} else {
				result[i] = boolFalse
			}
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 9: 仅支持常用类型（激进） ====================
func GetStringSliceV9(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	switch x := val.(type) {
	case []string:
		return x
	case []interface{}:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	case []int64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 方案 10: 混合策略 ====================
func GetStringSliceV10(p *MapAny, key string) []string {
	val, ok := p.get(key)
	if !ok {
		return nil
	}
	if x, ok := val.([]string); ok {
		return x
	}
	if x, ok := val.([]interface{}); ok {
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = candy.ToString(x[i])
		}
		return result
	}
	switch x := val.(type) {
	case []int64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatInt(x[i], 10)
		}
		return result
	case []uint64:
		n := len(x)
		result := make([]string, n)
		for i := 0; i < n; i++ {
			result[i] = strconv.FormatUint(x[i], 10)
		}
		return result
	default:
		return []string{}
	}
}

// ==================== 基准测试用例 ====================

// 测试 []string 类型
func BenchmarkGetStringSlice_Original_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceOrig(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V1_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV1(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V2_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV2(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V3_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV3(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V4_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV4(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V5_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV5(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V6_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV6(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V7_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV7(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V8_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV8(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V9_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV9(m, "string_slice")
	}
}

func BenchmarkGetStringSlice_V10_String(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV10(m, "string_slice")
	}
}

// 测试 []interface{} 类型
func BenchmarkGetStringSlice_Original_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceOrig(m, "interface_slice")
	}
}

func BenchmarkGetStringSlice_V1_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV1(m, "interface_slice")
	}
}

func BenchmarkGetStringSlice_V4_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV4(m, "interface_slice")
	}
}

func BenchmarkGetStringSlice_V5_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV5(m, "interface_slice")
	}
}

func BenchmarkGetStringSlice_V6_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV6(m, "interface_slice")
	}
}

func BenchmarkGetStringSlice_V8_Interface(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV8(m, "interface_slice")
	}
}

// 测试 []int64 类型
func BenchmarkGetStringSlice_Original_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceOrig(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V1_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV1(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V2_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV2(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V3_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV3(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V5_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV5(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V7_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV7(m, "int64_slice")
	}
}

func BenchmarkGetStringSlice_V8_Int64(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV8(m, "int64_slice")
	}
}

// 测试 []bool 类型
func BenchmarkGetStringSlice_Original_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceOrig(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V1_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV1(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V2_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV2(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V3_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV3(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V5_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV5(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V7_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV7(m, "bool_slice")
	}
}

func BenchmarkGetStringSlice_V8_Bool(b *testing.B) {
	m := setupBenchMap()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetStringSliceV8(m, "bool_slice")
	}
}
