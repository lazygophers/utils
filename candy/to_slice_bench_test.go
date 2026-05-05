package candy

import (
	"testing"
)

// 基准测试：ToInterfaceSlice 不同实现方案的性能对比

// 方案1：当前实现（baseline）- 使用 append
func ToInterfaceSliceBaseline(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []bool:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int8:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int16:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []int64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint8:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint16:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []uint64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []float32:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []float64:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []string:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case [][]byte:
		var v []interface{}
		for _, val := range x {
			v = append(v, val)
		}
		return v
	case []interface{}:
		return x
	default:
		return []interface{}{}
	}
}

// 方案2：预分配容量 + 直接索引
func ToInterfaceSlicePreallocated(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []bool:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int8:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int16:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []int64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint8:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint16:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []uint64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []float32:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []float64:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []string:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case [][]byte:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []interface{}:
		return x
	default:
		return []interface{}{}
	}
}

// 方案3：针对 []interface{} 的零拷贝优化
func ToInterfaceSliceZeroCopy(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []interface{}:
		return x
	case []int:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	case []string:
		v := make([]interface{}, len(x))
		for i := range x {
			v[i] = x[i]
		}
		return v
	default:
		return ToInterfaceSlicePreallocated(val)
	}
}

// 方案4：批量处理（每次处理4个元素）
func ToInterfaceSliceBatch4(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []interface{}:
		return x
	case []int:
		v := make([]interface{}, len(x))
		i := 0
		for ; i < len(x)-3; i += 4 {
			v[i] = x[i]
			v[i+1] = x[i+1]
			v[i+2] = x[i+2]
			v[i+3] = x[i+3]
		}
		for ; i < len(x); i++ {
			v[i] = x[i]
		}
		return v
	default:
		return ToInterfaceSlicePreallocated(val)
	}
}

// 方案5：批量处理（每次处理8个元素）
func ToInterfaceSliceBatch8(val interface{}) []interface{} {
	if val == nil {
		return nil
	}
	switch x := val.(type) {
	case []interface{}:
		return x
	case []int:
		v := make([]interface{}, len(x))
		i := 0
		for ; i < len(x)-7; i += 8 {
			v[i] = x[i]
			v[i+1] = x[i+1]
			v[i+2] = x[i+2]
			v[i+3] = x[i+3]
			v[i+4] = x[i+4]
			v[i+5] = x[i+5]
			v[i+6] = x[i+6]
			v[i+7] = x[i+7]
		}
		for ; i < len(x); i++ {
			v[i] = x[i]
		}
		return v
	default:
		return ToInterfaceSlicePreallocated(val)
	}
}

// 测试数据准备
func makeIntSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func makeStringSlice(n int) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = "test"
	}
	return s
}

func makeInterfaceSlice(n int) []interface{} {
	s := make([]interface{}, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

// 小数据集测试（10个元素）
func BenchmarkToInterfaceSlice_Small_Int_Baseline(b *testing.B) {
	data := makeIntSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_Small_Int_Preallocated(b *testing.B) {
	data := makeIntSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSlicePreallocated(data)
	}
}

func BenchmarkToInterfaceSlice_Small_Int_ZeroCopy(b *testing.B) {
	data := makeIntSlice(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceZeroCopy(data)
	}
}

// 中等数据集测试（100个元素）
func BenchmarkToInterfaceSlice_Medium_Int_Baseline(b *testing.B) {
	data := makeIntSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_Medium_Int_Preallocated(b *testing.B) {
	data := makeIntSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSlicePreallocated(data)
	}
}

func BenchmarkToInterfaceSlice_Medium_Int_Batch4(b *testing.B) {
	data := makeIntSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch4(data)
	}
}

func BenchmarkToInterfaceSlice_Medium_Int_Batch8(b *testing.B) {
	data := makeIntSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch8(data)
	}
}

// 大数据集测试（1000个元素）
func BenchmarkToInterfaceSlice_Large_Int_Baseline(b *testing.B) {
	data := makeIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_Large_Int_Preallocated(b *testing.B) {
	data := makeIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSlicePreallocated(data)
	}
}

func BenchmarkToInterfaceSlice_Large_Int_Batch4(b *testing.B) {
	data := makeIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch4(data)
	}
}

func BenchmarkToInterfaceSlice_Large_Int_Batch8(b *testing.B) {
	data := makeIntSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch8(data)
	}
}

// string 类型测试
func BenchmarkToInterfaceSlice_Medium_String_Baseline(b *testing.B) {
	data := makeStringSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_Medium_String_Preallocated(b *testing.B) {
	data := makeStringSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSlicePreallocated(data)
	}
}

// []interface{} 类型测试（零拷贝优化）
func BenchmarkToInterfaceSlice_Medium_Interface_Baseline(b *testing.B) {
	data := makeInterfaceSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_Medium_Interface_ZeroCopy(b *testing.B) {
	data := makeInterfaceSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceZeroCopy(data)
	}
}

// 超大数据集测试（10000个元素）
func BenchmarkToInterfaceSlice_XLarge_Int_Baseline(b *testing.B) {
	data := makeIntSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBaseline(data)
	}
}

func BenchmarkToInterfaceSlice_XLarge_Int_Preallocated(b *testing.B) {
	data := makeIntSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSlicePreallocated(data)
	}
}

func BenchmarkToInterfaceSlice_XLarge_Int_Batch4(b *testing.B) {
	data := makeIntSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch4(data)
	}
}

func BenchmarkToInterfaceSlice_XLarge_Int_Batch8(b *testing.B) {
	data := makeIntSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ToInterfaceSliceBatch8(data)
	}
}
