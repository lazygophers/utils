package candy

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)




func All[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if !f(s) {
			return false
		}
	}

	return true
}

func Shuffle[T any](ss []T) []T {
	for i := range ss {
		j := rand.Intn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}

	return ss
}

func Max[T constraints.Ordered](ss []T) (max T) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}

func Min[T constraints.Ordered](ss []T) (min T) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}

func Unique[T constraints.Ordered](ss []T) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0)
	m := make(map[T]struct{}, len(ss))
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}

	return
}

func UniqueUsing[T any](ss []T, f func(T) any) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0)
	m := make(map[any]T, len(ss))
	for _, s := range ss {
		if _, ok := m[(f(s))]; !ok {
			m[(f(s))] = s
			ret = append(ret, s)
		}
	}

	return
}

func Random[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[rand.Intn(len(ss))]
}

func Each[T any](ss []T, f func(T)) {
	for _, s := range ss {
		f(s)
	}
}

func EachStopWithError[T any](ss []T, f func(T) (err error)) (err error) {
	for _, s := range ss {
		err = f(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func EachReverse[T any](ss []T, f func(T)) {
	for i := len(ss) - 1; i >= 0; i-- {
		f(ss[i])
	}
}

func Sort[T constraints.Ordered](ss []T) []T {
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

func SortUsing[T any](ss []T, less func(a, b T) bool) []T {
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]T, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})

	return sorted
}

func Filter[T any](ss []T, f func(T) bool) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0)
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return
}

func Map[T, U any](ss []T, f func(T) U) (ret []U) {
	ret = make([]U, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, f(s))
	}

	return
}

func Contains[T constraints.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}

func ContainsUsing[T any](ss []T, f func(v T) bool) bool {
	for _, v := range ss {
		if f(v) {
			return true
		}
	}

	return false
}

func Sum[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	for _, s := range ss {
		ret += s
	}

	return
}

func First[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[0]
}

func FirstOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[0]
}

func Last[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[len(ss)-1]
}

func LastOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[len(ss)-1]
}

func Top[T any](ss []T, n int) (ret []T) {
	if n > len(ss) {
		n = len(ss)
	}

	ret = make([]T, n)
	copy(ret, ss[:n])
	return ret
}

func Bottom[T any](ss []T, n int) (ret []T) {
	if n > len(ss) {
		n = len(ss)
	}

	ret = make([]T, n)
	copy(ret, ss[len(ss)-n:])
	return ret
}

func Average[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	// 使用 float64 计算以保持精度
	var sum float64
	for _, s := range ss {
		sum += float64(s)
	}
	return T(sum / float64(len(ss)))
}

func Reverse[T any](ss []T) (ret []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	ret = make([]T, 0, len(ss))
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}

	return
}

func Chunk[T any](ss []T, size int) (ret [][]T) {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}

	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}

		ret = append(ret, ss[i:end])
	}

	return
}

func Index[T constraints.Ordered](ss []T, sub T) int {
	if len(ss) == 0 {
		return -1
	}

	for i, s := range ss {
		if s == sub {
			return i
		}
	}

	return -1
}

// Diff
//
//	added: ss 不存在 against 存在
//	removed: ss 存在 against 不存在
func Diff[T constraints.Ordered](ss []T, against []T) (added, removed []T) {
	removed = Remove(ss, against)
	added = Remove(against, ss)

	return
}

// Remove 移除 ss 存在也 against 存在的部分
// 返回在 against 中但不在 ss 中的元素
func Remove[T constraints.Ordered](ss []T, against []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; !ok {
			result = append(result, s)
		}
	}
	return
}

// Same 两个同事存在的
// 返回在 against 和 ss 中都存在的元素
func Same[T constraints.Ordered](against []T, ss []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return
}

// Spare ss 不存在但是 against 存在
// 返回在 against 中但不在 ss 中的元素（与 Remove 功能相同）
func Spare[T constraints.Ordered](ss []T, against []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; !ok {
			result = append(result, s)
		}
	}
	return
}

// RemoveIndex 移除指定索引的元素
func RemoveIndex[T any](ss []T, index int) []T {
	if len(ss) == 0 || index < 0 || index >= len(ss) {
		// 返回空切片而非原切片
		return make([]T, 0)
	}

	if index == 0 {
		if len(ss) > 0 {
			return ss[1:]
		}

		return make([]T, 0)
	}

	if index == len(ss)-1 {
		return ss[:len(ss)-1]
	}

	return append(ss[:index], ss[index+1:]...)
}

func SliceEqual[T any](a, b []T) bool {
	// 处理 nil 切片的情况：nil 和空切片视为相等
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	// 使用 map 来跟踪每个元素的出现次数
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

	// 检查所有元素的计数是否都为0
	for _, count := range am {
		if count != 0 {
			return false
		}
	}

	return true
}

func String[T constraints.Ordered](s T) string {
	return fmt.Sprintf("%v", s)
}

func Join[T constraints.Ordered](ss []T, glue ...string) string {
	seq := ","
	if len(glue) > 0 {
		seq = glue[0]
	}

	return strings.Join(Map(ss, func(s T) string {
		return String(s)
	}), seq)
}
