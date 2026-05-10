package anyx

import (
	"strconv"
	"testing"
)

// =====================================================
// GetInt 性能优化 Benchmark
// =====================================================
// 目标：优化 GetInt 函数性能，设计不少于 10 种方案进行对比测试
// =====================================================

// 方案 1: 当前实现 - 调用 candy.ToInt
func getMethodImpl1(val interface{}) int {
	switch v := val.(type) {
	case nil:
		return 0
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 2: 快速路径优化 - 最常见类型优先
func getMethodImpl2(val interface{}) int {
	// 快速路径：最常见类型优先
	if v, ok := val.(int); ok {
		return v
	}
	if v, ok := val.(int64); ok {
		return int(v)
	}
	if v, ok := val.(float64); ok {
		return int(v)
	}
	if v, ok := val.(string); ok {
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	}
	if val == nil {
		return 0
	}

	// 慢速路径：其他类型
	switch v := val.(type) {
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 3: 内联所有逻辑 - 避免任何函数调用
func getMethodImpl3(val interface{}) int {
	switch v := val.(type) {
	case nil:
		return 0
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 4: 整数类型合并处理
func getMethodImpl4(val interface{}) int {
	if val == nil {
		return 0
	}

	// 整数类型统一处理
	switch v := val.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	}

	// 浮点类型
	switch v := val.(type) {
	case float32:
		return int(v)
	case float64:
		return int(v)
	}

	// 字符串类型
	switch v := val.(type) {
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	}

	// 布尔类型
	if v, ok := val.(bool); ok {
		if v {
			return 1
		}
		return 0
	}

	return 0
}

// 方案 5: 零拷贝优化 - 特定类型直接返回
func getMethodImpl5(val interface{}) int {
	// 零拷贝路径：相同类型直接返回
	if v, ok := val.(int); ok {
		return v
	}

	// 其他路径需要类型转换
	if val == nil {
		return 0
	}

	switch v := val.(type) {
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 6: 分支预测优化 - 热路径优先
func getMethodImpl6(val interface{}) int {
	// 分支预测优化：按概率排序
	switch v := val.(type) {
	case int:
		return v
	case string:
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case float64:
		return int(v)
	case nil:
		return 0
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 7: 内联字符串解析 - 避免函数调用
func getMethodImpl7(val interface{}) int {
	switch v := val.(type) {
	case nil:
		return 0
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		// 内联字符串解析逻辑
		if v == "" {
			return 0
		}
		neg := false
		if v[0] == '-' {
			neg = true
			v = v[1:]
		} else if v[0] == '+' {
			v = v[1:]
		}
		result := int64(0)
		for _, c := range v {
			if c < '0' || c > '9' {
				return 0
			}
			result = result*10 + int64(c-'0')
		}
		if neg {
			result = -result
		}
		return int(result)
	case []byte:
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// 方案 8: 最小化类型断言 - 统一处理数值类型
func getMethodImpl8(val interface{}) int {
	// 快速路径：int 类型
	if v, ok := val.(int); ok {
		return v
	}

	// nil 检查
	if val == nil {
		return 0
	}

	// 字符串类型处理
	switch v := val.(type) {
	case string:
		if num, err := strconv.ParseInt(v, 10, 0); err == nil {
			return int(num)
		}
		return 0
	case []byte:
		if num, err := strconv.ParseInt(string(v), 10, 0); err == nil {
			return int(num)
		}
		return 0
	}

	// 数值类型统一处理
	switch v := val.(type) {
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// 方案 9: 最小化分支 - 短路求值
func getMethodImpl9(val interface{}) int {
	// 短路求值：nil 优先检查
	if val == nil {
		return 0
	}

	// 快速路径：int 类型
	if i, ok := val.(int); ok {
		return i
	}

	// 统一处理其他类型
	switch v := val.(type) {
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 0); err == nil {
			return int(i)
		}
		return 0
	case []byte:
		if i, err := strconv.ParseInt(string(v), 10, 0); err == nil {
			return int(i)
		}
		return 0
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// 方案 10: 类型分层处理 - 渐进式优化
func getMethodImpl10(val interface{}) int {
	// 第一层：零成本转换
	if v, ok := val.(int); ok {
		return v
	}

	// 第二层：低成本转换
	if v, ok := val.(int64); ok {
		return int(v)
	}

	// 第三层：中等成本转换
	switch v := val.(type) {
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		if i, err := strconv.ParseInt(v, 10, 0); err == nil {
			return int(i)
		}
		return 0
	case []byte:
		if i, err := strconv.ParseInt(string(v), 10, 0); err == nil {
			return int(i)
		}
		return 0
	case bool:
		if v {
			return 1
		}
		return 0
	case nil:
		return 0
	default:
		return 0
	}
}

// 方案 11: 组合优化 - 零拷贝 + 快速路径 + 分支预测
func getMethodImpl11(val interface{}) int {
	// 零拷贝快速路径
	if v, ok := val.(int); ok {
		return v
	}

	// nil 快速检查
	if val == nil {
		return 0
	}

	// 分支预测优化：按热度和转换成本排序
	switch v := val.(type) {
	case int64: // 常见且转换成本低
		return int(v)
	case float64: // 常见但转换成本中等
		return int(v)
	case string: // 常见但转换成本高
		val, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	case int8: // 较少见但转换成本低
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case uint: // 较少见但转换成本低
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32: // 较少见且转换成本中等
		return int(v)
	case bool: // 较少见
		if v {
			return 1
		}
		return 0
	case []byte: // 少见且转换成本高
		val, err := strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			return 0
		}
		return int(val)
	default:
		return 0
	}
}

// =====================================================
// Benchmark 测试用例
// =====================================================

// BenchmarkGetInt_AllImplementations_Int - 测试所有实现的 int 类型性能
func BenchmarkGetInt_AllImplementations_Int(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl2_FastPath", getMethodImpl2},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl4_IntGrouped", getMethodImpl4},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl6_BranchPred", getMethodImpl6},
		{"Impl7_InlineParse", getMethodImpl7},
		{"Impl8_MinAssert", getMethodImpl8},
		{"Impl9_MinBranch", getMethodImpl9},
		{"Impl10_Layered", getMethodImpl10},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}(42)

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

// BenchmarkGetInt_AllImplementations_String - 测试所有实现的 string 类型性能
func BenchmarkGetInt_AllImplementations_String(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl2_FastPath", getMethodImpl2},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl4_IntGrouped", getMethodImpl4},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl6_BranchPred", getMethodImpl6},
		{"Impl7_InlineParse", getMethodImpl7},
		{"Impl8_MinAssert", getMethodImpl8},
		{"Impl9_MinBranch", getMethodImpl9},
		{"Impl10_Layered", getMethodImpl10},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}("123")

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

// BenchmarkGetInt_AllImplementations_Float64 - 测试所有实现的 float64 类型性能
func BenchmarkGetInt_AllImplementations_Float64(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl2_FastPath", getMethodImpl2},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl4_IntGrouped", getMethodImpl4},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl6_BranchPred", getMethodImpl6},
		{"Impl7_InlineParse", getMethodImpl7},
		{"Impl8_MinAssert", getMethodImpl8},
		{"Impl9_MinBranch", getMethodImpl9},
		{"Impl10_Layered", getMethodImpl10},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}(42.5)

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

// BenchmarkGetInt_AllImplementations_Int64 - 测试所有实现的 int64 类型性能
func BenchmarkGetInt_AllImplementations_Int64(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl2_FastPath", getMethodImpl2},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl4_IntGrouped", getMethodImpl4},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl6_BranchPred", getMethodImpl6},
		{"Impl7_InlineParse", getMethodImpl7},
		{"Impl8_MinAssert", getMethodImpl8},
		{"Impl9_MinBranch", getMethodImpl9},
		{"Impl10_Layered", getMethodImpl10},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}(int64(42))

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

// BenchmarkGetInt_AllImplementations_Nil - 测试所有实现的 nil 类型性能
func BenchmarkGetInt_AllImplementations_Nil(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl2_FastPath", getMethodImpl2},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl4_IntGrouped", getMethodImpl4},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl6_BranchPred", getMethodImpl6},
		{"Impl7_InlineParse", getMethodImpl7},
		{"Impl8_MinAssert", getMethodImpl8},
		{"Impl9_MinBranch", getMethodImpl9},
		{"Impl10_Layered", getMethodImpl10},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}(nil)

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

// 内存分配对比测试
func BenchmarkGetInt_Allocation_Int(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}(42)

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}

func BenchmarkGetInt_Allocation_String(b *testing.B) {
	impls := []struct {
		name string
		fn   func(interface{}) int
	}{
		{"Impl1_Current", getMethodImpl1},
		{"Impl3_Inlined", getMethodImpl3},
		{"Impl5_ZeroCopy", getMethodImpl5},
		{"Impl11_Combined", getMethodImpl11},
	}

	testVal := interface{}("123")

	for _, impl := range impls {
		b.Run(impl.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = impl.fn(testVal)
			}
		})
	}
}
