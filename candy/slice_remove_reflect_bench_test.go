package candy

import (
	"reflect"
	"testing"
)

// Benchmark RemoveSlice 反射函数的各种优化方案

// 方案1: 当前实现（快速路径 + 反射）
func BenchmarkRemoveSlice_Current_Int_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSlice(data, toRemove)
	}
}

// 方案2: 纯反射（无快速路径）
func RemoveSlicePureReflect(src interface{}, rm interface{}) interface{} {
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("src is not slice")
	}

	bt := reflect.TypeOf(rm)
	if bt.Kind() != reflect.Slice {
		panic("rm is not slice")
	}

	if at.Elem().Kind() != bt.Elem().Kind() {
		panic("src and rm are not same type")
	}

	m := map[interface{}]bool{}
	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}

	av := reflect.ValueOf(src)
	c := reflect.MakeSlice(at, 0, av.Len()-bv.Len()/2)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
		}
	}

	return c.Interface()
}

func BenchmarkRemoveSlice_PureReflect_Int_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSlicePureReflect(data, toRemove)
	}
}

// 方案3: 优化的快速路径（使用索引循环）
func RemoveSliceOptFastPath(src interface{}, rm interface{}) interface{} {
	// 快速路径：int 类型
	if srcInt, ok := src.([]int); ok {
		if rmInt, ok := rm.([]int); ok {
			m := make(map[int]struct{}, len(rmInt))
			for i := 0; i < len(rmInt); i++ {
				m[rmInt[i]] = struct{}{}
			}
			result := make([]int, 0, len(srcInt)-len(rmInt)/2)
			for i := 0; i < len(srcInt); i++ {
				if _, ok := m[srcInt[i]]; !ok {
					result = append(result, srcInt[i])
				}
			}
			return result
		}
	}

	// 快速路径：string 类型
	if srcStr, ok := src.([]string); ok {
		if rmStr, ok := rm.([]string); ok {
			m := make(map[string]struct{}, len(rmStr))
			for i := 0; i < len(rmStr); i++ {
				m[rmStr[i]] = struct{}{}
			}
			result := make([]string, 0, len(srcStr)-len(rmStr)/2)
			for i := 0; i < len(srcStr); i++ {
				if _, ok := m[srcStr[i]]; !ok {
					result = append(result, srcStr[i])
				}
			}
			return result
		}
	}

	// 反射路径
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("src is not slice")
	}

	bt := reflect.TypeOf(rm)
	if bt.Kind() != reflect.Slice {
		panic("rm is not slice")
	}

	if at.Elem().Kind() != bt.Elem().Kind() {
		panic("src and rm are not same type")
	}

	m := map[interface{}]bool{}
	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}

	av := reflect.ValueOf(src)
	c := reflect.MakeSlice(at, 0, av.Len()-bv.Len()/2)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
		}
	}

	return c.Interface()
}

func BenchmarkRemoveSlice_OptFastPath_Int_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceOptFastPath(data, toRemove)
	}
}

// 方案4: 使用 bool 代替 struct{}（快速路径）
func RemoveSliceBoolMap(src interface{}, rm interface{}) interface{} {
	// 快速路径：int 类型
	if srcInt, ok := src.([]int); ok {
		if rmInt, ok := rm.([]int); ok {
			m := make(map[int]bool, len(rmInt))
			for i := 0; i < len(rmInt); i++ {
				m[rmInt[i]] = true
			}
			result := make([]int, 0, len(srcInt)-len(rmInt)/2)
			for i := 0; i < len(srcInt); i++ {
				if !m[srcInt[i]] {
					result = append(result, srcInt[i])
				}
			}
			return result
		}
	}

	// 快速路径：string 类型
	if srcStr, ok := src.([]string); ok {
		if rmStr, ok := rm.([]string); ok {
			m := make(map[string]bool, len(rmStr))
			for i := 0; i < len(rmStr); i++ {
				m[rmStr[i]] = true
			}
			result := make([]string, 0, len(srcStr)-len(rmStr)/2)
			for i := 0; i < len(srcStr); i++ {
				if !m[srcStr[i]] {
					result = append(result, srcStr[i])
				}
			}
			return result
		}
	}

	// 反射路径
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("src is not slice")
	}

	bt := reflect.TypeOf(rm)
	if bt.Kind() != reflect.Slice {
		panic("rm is not slice")
	}

	if at.Elem().Kind() != bt.Elem().Kind() {
		panic("src and rm are not same type")
	}

	m := map[interface{}]bool{}
	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}

	av := reflect.ValueOf(src)
	c := reflect.MakeSlice(at, 0, av.Len()-bv.Len()/2)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
		}
	}

	return c.Interface()
}

func BenchmarkRemoveSlice_BoolMap_Int_Small(b *testing.B) {
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = i
	}
	toRemove := []int{10, 20, 30, 40, 50}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceBoolMap(data, toRemove)
	}
}

// 中等数据集测试
func BenchmarkRemoveSlice_Current_Int_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSlice(data, toRemove)
	}
}

func BenchmarkRemoveSlice_PureReflect_Int_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSlicePureReflect(data, toRemove)
	}
}

func BenchmarkRemoveSlice_OptFastPath_Int_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceOptFastPath(data, toRemove)
	}
}

func BenchmarkRemoveSlice_BoolMap_Int_Medium(b *testing.B) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	toRemove := make([]int, 100)
	for i := 0; i < 100; i++ {
		toRemove[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceBoolMap(data, toRemove)
	}
}

// String 类型测试
func BenchmarkRemoveSlice_Current_String_Small(b *testing.B) {
	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = string(rune('a' + i))
	}
	toRemove := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSlice(data, toRemove)
	}
}

func BenchmarkRemoveSlice_OptFastPath_String_Small(b *testing.B) {
	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = string(rune('a' + i))
	}
	toRemove := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceOptFastPath(data, toRemove)
	}
}

func BenchmarkRemoveSlice_BoolMap_String_Small(b *testing.B) {
	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = string(rune('a' + i))
	}
	toRemove := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RemoveSliceBoolMap(data, toRemove)
	}
}
