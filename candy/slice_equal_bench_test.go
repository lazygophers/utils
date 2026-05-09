package candy

import (
	"reflect"
	"slices"
	"testing"

	"golang.org/x/exp/constraints"
)

// ========== Diff 函数优化方案 ==========

// DiffOriginal - 原始实现（使用 Remove）
func DiffOriginal[T constraints.Ordered](ss1, ss2 []T) []T {
	removed, _ := Diff(ss1, ss2)
	added, _ := Diff(ss2, ss1)
	result := make([]T, 0, len(removed)+len(added))
	result = append(result, removed...)
	result = append(result, added...)
	return result
}

// DiffOpt1 - 使用单一 map，避免两次 Remove 调用
func DiffOpt1[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	// 构建两个 map
	map1 := make(map[T]int, len(ss1))
	for _, v := range ss1 {
		map1[v]++
	}

	map2 := make(map[T]int, len(ss2))
	for _, v := range ss2 {
		map2[v]++
	}

	result := []T{}

	// 找出在 ss1 中但不在 ss2 中的
	for k, v := range map1 {
		if map2[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	// 找出在 ss2 中但不在 ss1 中的
	for k, v := range map2 {
		if map1[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt2 - 预分配结果切片
func DiffOpt2[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	map1 := make(map[T]int, len(ss1))
	for _, v := range ss1 {
		map1[v]++
	}

	map2 := make(map[T]int, len(ss2))
	for _, v := range ss2 {
		map2[v]++
	}

	// 预分配：最大可能大小是 len(ss1) + len(ss2)
	result := make([]T, 0, len(ss1)+len(ss2))

	for k, v := range map1 {
		if map2[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	for k, v := range map2 {
		if map1[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt3 - 使用 struct{} 代替 int 作为 map 值（仅用于判断存在性）
func DiffOpt3[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	set1 := make(map[T]struct{}, len(ss1))
	for _, v := range ss1 {
		set1[v] = struct{}{}
	}

	set2 := make(map[T]struct{}, len(ss2))
	for _, v := range ss2 {
		set2[v] = struct{}{}
	}

	result := make([]T, 0, len(set1)+len(set2))

	for k := range set1 {
		if _, ok := set2[k]; !ok {
			result = append(result, k)
		}
	}

	for k := range set2 {
		if _, ok := set1[k]; !ok {
			result = append(result, k)
		}
	}

	return result
}

// DiffOpt4 - 先统计差异数量，再精确分配
func DiffOpt4[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	map1 := make(map[T]int, len(ss1))
	for _, v := range ss1 {
		map1[v]++
	}

	map2 := make(map[T]int, len(ss2))
	for _, v := range ss2 {
		map2[v]++
	}

	// 第一遍：计算差异总数
	count := 0
	for k, v := range map1 {
		if map2[k] == 0 {
			count += v
		}
	}
	for k, v := range map2 {
		if map1[k] == 0 {
			count += v
		}
	}

	// 精确分配
	result := make([]T, 0, count)
	for k, v := range map1 {
		if map2[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}
	for k, v := range map2 {
		if map1[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt5 - 针对 int 类型的特化优化
func DiffOpt5Int(ss1, ss2 []int) []int {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []int{}
	}

	map1 := make(map[int]int, len(ss1))
	for _, v := range ss1 {
		map1[v]++
	}

	map2 := make(map[int]int, len(ss2))
	for _, v := range ss2 {
		map2[v]++
	}

	result := make([]int, 0, len(map1)+len(map2))

	for k, v := range map1 {
		if map2[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	for k, v := range map2 {
		if map1[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt6 - 使用 slices.Contains 进行小切片优化
func DiffOpt6[T constraints.Ordered](ss1, ss2 []T) []T {
	const smallSliceThreshold = 32

	if len(ss1) < smallSliceThreshold && len(ss2) < smallSliceThreshold {
		result := []T{}
		for _, v := range ss1 {
			if !slices.Contains(ss2, v) {
				result = append(result, v)
			}
		}
		for _, v := range ss2 {
			if !slices.Contains(ss1, v) {
				result = append(result, v)
			}
		}
		return result
	}

	return DiffOpt2(ss1, ss2)
}

// DiffOpt7 - 使用 reflect.DeepEqual 比较（仅用于兼容性）
func DiffOpt7[T constraints.Ordered](ss1, ss2 []T) []T {
	if reflect.DeepEqual(ss1, ss2) {
		return []T{}
	}
	return DiffOpt2(ss1, ss2)
}

// DiffOpt8 - 延迟 map 创建
func DiffOpt8[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 {
		result := make([]T, 0, len(ss2))
		return append(result, ss2...)
	}
	if len(ss2) == 0 {
		result := make([]T, 0, len(ss1))
		return append(result, ss1...)
	}

	map1 := make(map[T]int, len(ss1))
	for _, v := range ss1 {
		map1[v]++
	}

	map2 := make(map[T]int, len(ss2))
	for _, v := range ss2 {
		map2[v]++
	}

	result := make([]T, 0, len(map1)+len(map2))

	for k, v := range map1 {
		if map2[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	for k, v := range map2 {
		if map1[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt9 - 使用单个 map + 计数
func DiffOpt9[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	combined := make(map[T]int, len(ss1)+len(ss2))

	for _, v := range ss1 {
		combined[v]++
	}

	for _, v := range ss2 {
		combined[v]--
	}

	result := make([]T, 0, len(combined))
	for k, v := range combined {
		if v != 0 {
			abs := v
			if abs < 0 {
				abs = -abs
			}
			for i := 0; i < abs; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// DiffOpt10 - 批量处理
func DiffOpt10[T constraints.Ordered](ss1, ss2 []T) []T {
	if len(ss1) == 0 && len(ss2) == 0 {
		return []T{}
	}

	// 使用较小的 map
	smaller, larger := ss1, ss2
	if len(ss1) > len(ss2) {
		smaller, larger = ss2, ss1
	}

	smallerMap := make(map[T]int, len(smaller))
	for _, v := range smaller {
		smallerMap[v]++
	}

	largerMap := make(map[T]int, len(larger))
	for _, v := range larger {
		largerMap[v]++
	}

	result := make([]T, 0, len(smallerMap)+len(largerMap))

	for k, v := range smallerMap {
		if largerMap[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	for k, v := range largerMap {
		if smallerMap[k] == 0 {
			for i := 0; i < v; i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

// ========== Same 函数优化方案 ==========

// SameOriginal - 原始实现
func SameOriginal[T constraints.Ordered](against, ss []T) []T {
	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	result := make([]T, 0)
	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt1 - 预分配结果切片
func SameOpt1[T constraints.Ordered](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 预分配：最大可能大小是 min(len(against), len(ss))
	maxSize := len(against)
	if len(ss) < maxSize {
		maxSize = len(ss)
	}

	result := make([]T, 0, maxSize)
	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt2 - 使用 map[T]bool 替代 struct{}
func SameOpt2[T constraints.Ordered](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	set := make(map[T]bool, len(ss))
	for _, s := range ss {
		set[s] = true
	}

	result := make([]T, 0, len(against))
	for _, s := range against {
		if set[s] {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt3 - 先计数交集元素
func SameOpt3[T constraints.Ordered](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 第一遍：计算交集大小
	count := 0
	for _, s := range against {
		if _, ok := set[s]; ok {
			count++
		}
	}

	result := make([]T, 0, count)
	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt4 - 小切片使用 slices.Contains
func SameOpt4[T constraints.Ordered](against, ss []T) []T {
	const smallSliceThreshold = 32

	if len(ss) < smallSliceThreshold {
		result := make([]T, 0, len(against))
		for _, s := range against {
			if slices.Contains(ss, s) {
				result = append(result, s)
			}
		}
		return result
	}

	return SameOpt1(against, ss)
}

// SameOpt5 - 针对 int 的特化
func SameOpt5Int(against, ss []int) []int {
	if len(ss) == 0 || len(against) == 0 {
		return []int{}
	}

	set := make(map[int]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	result := make([]int, 0, len(against))
	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt6 - 使用较小切片构建 map
func SameOpt6[T constraints.Ordered](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	// 使用较小的切片构建 map
	smaller, larger := ss, against
	if len(ss) > len(against) {
		smaller, larger = against, ss
	}

	set := make(map[T]struct{}, len(smaller))
	for _, s := range smaller {
		set[s] = struct{}{}
	}

	result := make([]T, 0, len(larger))
	for _, s := range larger {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt7 - 使用 int 计数而非 struct{}
func SameOpt7[T comparable](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	counts := make(map[T]int, len(ss))
	for _, s := range ss {
		counts[s]++
	}

	result := make([]T, 0, len(against))
	for _, s := range against {
		if counts[s] > 0 {
			result = append(result, s)
			counts[s]--
		}
	}
	return result
}

// SameOpt8 - 使用 slices.Equal 优化空切片
func SameOpt8[T constraints.Ordered](against, ss []T) []T {
	if slices.Equal(against, ss) {
		result := make([]T, len(against))
		copy(result, against)
		return result
	}

	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	result := make([]T, 0, len(against))
	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return result
}

// SameOpt9 - 使用 reflect.DeepEqual 预检查
func SameOpt9[T constraints.Ordered](against, ss []T) []T {
	if reflect.DeepEqual(against, ss) {
		result := make([]T, len(against))
		copy(result, against)
		return result
	}

	return SameOpt1(against, ss)
}

// SameOpt10 - 两阶段处理
func SameOpt10[T constraints.Ordered](against, ss []T) []T {
	if len(ss) == 0 || len(against) == 0 {
		return []T{}
	}

	// 第一阶段：快速跳过明显不存在的
	set := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 第二阶段：收集结果
	intersect := make([]T, 0, len(against))
	for _, s := range against {
		if _, exists := set[s]; exists {
			intersect = append(intersect, s)
		}
	}

	return intersect
}

// ========== Equal 函数优化方案 ==========

// EqualOriginal - 原始实现
func EqualOriginal[T comparable](a, b T) bool {
	return a == b
}

// EqualOpt1 - 使用 reflect.DeepEqual
func EqualOpt1[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}

// EqualOpt2 - 针对指针的优化
func EqualOpt2[T comparable](a, b T) bool {
	return a == b
}

// EqualOpt3 - 针对字符串的特化
func EqualOpt3String(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return a == b
}

// EqualOpt4 - 针对 int 的特化
func EqualOpt4Int(a, b int) bool {
	return a == b
}

// EqualOpt5 - 针对 float64 的特化
func EqualOpt5Float64(a, b float64) bool {
	return a == b
}

// EqualOpt6 - 使用 unsafe 指针比较（危险，仅测试用）
func EqualOpt6[T comparable](a, b T) bool {
	return a == b
}

// EqualOpt7 - 使用接口转换
func EqualOpt7[T comparable](a, b T) bool {
	ai := any(a)
	bi := any(b)
	return ai == bi
}

// EqualOpt8 - 针对 byte 的特化
func EqualOpt8Byte(a, b byte) bool {
	return a == b
}

// EqualOpt9 - 针对 rune 的特化
func EqualOpt9Rune(a, b rune) bool {
	return a == b
}

// EqualOpt10 - 使用泛型类型断言
func EqualOpt10[T comparable](a, b T) bool {
	switch v := any(a).(type) {
	case int:
		return v == any(b).(int)
	case string:
		return v == any(b).(string)
	default:
		return a == b
	}
}

// ========== SliceEqual 函数优化方案 ==========

// SliceEqualOriginal - 原始实现
func SliceEqualOriginal[T any](a, b []T) bool {
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
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	return true
}

// SliceEqualOpt1 - 使用 reflect.DeepEqual
func SliceEqualOpt1[T any](a, b []T) bool {
	return reflect.DeepEqual(a, b)
}

// SliceEqualOpt2 - 预分配 map 容量
func SliceEqualOpt2[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	am := make(map[T]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	return true
}

// SliceEqualOpt3 - 使用 struct{} 代替 int
func SliceEqualOpt3[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	set := make(map[T]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := set[v]; !ok {
			return false
		}
	}

	return true
}

// SliceEqualOpt4 - 先排序再比较
func SliceEqualOpt4[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]T, len(a))
	copy(aCopy, a)
	bCopy := make([]T, len(b))
	copy(bCopy, b)

	slices.Sort(aCopy)
	slices.Sort(bCopy)

	return slices.Equal(aCopy, bCopy)
}

// SliceEqualOpt5 - 针对小切片的线性扫描
func SliceEqualOpt5[T comparable](a, b []T) bool {
	const smallSliceThreshold = 32

	if len(a) != len(b) {
		return false
	}

	if len(a) < smallSliceThreshold {
		// 小切片使用双重循环
		for _, va := range a {
			found := false
			for _, vb := range b {
				if va == vb {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
		return true
	}

	// 大切片使用 map
	return SliceEqualOpt2(a, b)
}

// SliceEqualOpt6 - 针对 int 的特化
func SliceEqualOpt6Int(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	am := make(map[int]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	return true
}

// SliceEqualOpt7 - 针对 string 的特化
func SliceEqualOpt7String(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	am := make(map[string]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	return true
}

// SliceEqualOpt8 - 使用两个 map
func SliceEqualOpt8[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	mapA := make(map[T]int, len(a))
	for _, v := range a {
		mapA[v]++
	}

	mapB := make(map[T]int, len(b))
	for _, v := range b {
		mapB[v]++
	}

	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}

	return true
}

// SliceEqualOpt9 - 使用 slices.EqualSorted（需要先排序）
func SliceEqualOpt9[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// 检查是否已排序
	if slices.IsSorted(a) && slices.IsSorted(b) {
		return slices.Equal(a, b)
	}

	// 未排序则先排序再比较
	aCopy := slices.Clone(a)
	bCopy := slices.Clone(b)
	slices.Sort(aCopy)
	slices.Sort(bCopy)

	return slices.Equal(aCopy, bCopy)
}

// SliceEqualOpt10 - 使用单一 map + 计数
func SliceEqualOpt10[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	combined := make(map[T]int, len(a)+len(b))

	for _, v := range a {
		combined[v]++
	}

	for _, v := range b {
		combined[v]--
	}

	for _, v := range combined {
		if v != 0 {
			return false
		}
	}

	return true
}

// ========== 基准测试 ==========

// 测试数据
var (
	smallSliceInt  []int    = []int{1, 2, 3, 4, 5}
	mediumSliceInt []int    = make([]int, 100)
	largeSliceInt  []int    = make([]int, 1000)
	smallSliceStr  []string = []string{"a", "b", "c", "d", "e"}
	mediumSliceStr []string = make([]string, 100)
	largeSliceStr  []string = make([]string, 1000)
)

func init() {
	for i := 0; i < 100; i++ {
		mediumSliceInt[i] = i % 50
		mediumSliceStr[i] = string(rune('a' + i%26))
	}
	for i := 0; i < 1000; i++ {
		largeSliceInt[i] = i % 100
		largeSliceStr[i] = string(rune('a' + i%26))
	}
}

// Diff 基准测试
func BenchmarkDiffOriginal_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOriginal(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOpt1_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt1(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOpt2_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt2(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOpt3_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt3(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOpt9_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt9(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOpt10_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt10(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkDiffOriginal_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOriginal(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkDiffOpt1_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt1(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkDiffOpt2_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt2(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkDiffOpt9_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt9(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkDiffOpt10_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DiffOpt10(mediumSliceInt, mediumSliceInt[:50])
	}
}

// Same 基准测试
func BenchmarkSameOriginal_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOriginal(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkSameOpt1_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt1(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkSameOpt2_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt2(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkSameOpt3_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt3(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkSameOpt6_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt6(smallSliceInt, []int{3, 4, 5, 6})
	}
}

func BenchmarkSameOriginal_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOriginal(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkSameOpt1_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt1(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkSameOpt2_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt2(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkSameOpt3_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt3(mediumSliceInt, mediumSliceInt[:50])
	}
}

func BenchmarkSameOpt6_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SameOpt6(mediumSliceInt, mediumSliceInt[:50])
	}
}

// Equal 基准测试
func BenchmarkEqualOriginal_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualOriginal(42, 42)
	}
}

func BenchmarkEqualOpt1_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualOpt1(42, 42)
	}
}

func BenchmarkEqualOpt3_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualOpt3String("hello", "hello")
	}
}

func BenchmarkEqualOriginal_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualOriginal("hello", "hello")
	}
}

// SliceEqual 基准测试
func BenchmarkSliceEqualOriginal_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOriginal(smallSliceInt, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSliceEqualOpt1_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt1(smallSliceInt, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSliceEqualOpt2_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt2(smallSliceInt, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSliceEqualOpt5_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt5(smallSliceInt, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSliceEqualOpt6_Small(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt6Int(smallSliceInt, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSliceEqualOriginal_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOriginal(mediumSliceInt, mediumSliceInt)
	}
}

func BenchmarkSliceEqualOpt1_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt1(mediumSliceInt, mediumSliceInt)
	}
}

func BenchmarkSliceEqualOpt2_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt2(mediumSliceInt, mediumSliceInt)
	}
}

func BenchmarkSliceEqualOpt5_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt5(mediumSliceInt, mediumSliceInt)
	}
}

func BenchmarkSliceEqualOpt6_Medium(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt6Int(mediumSliceInt, mediumSliceInt)
	}
}

func BenchmarkSliceEqualOriginal_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOriginal(largeSliceInt, largeSliceInt)
	}
}

func BenchmarkSliceEqualOpt1_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt1(largeSliceInt, largeSliceInt)
	}
}

func BenchmarkSliceEqualOpt2_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt2(largeSliceInt, largeSliceInt)
	}
}

func BenchmarkSliceEqualOpt5_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt5(largeSliceInt, largeSliceInt)
	}
}

func BenchmarkSliceEqualOpt6_Large(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceEqualOpt6Int(largeSliceInt, largeSliceInt)
	}
}
