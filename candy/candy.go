package candy

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
	"sort"
)

func Abs[T constraints.Integer | constraints.Float](s T) T {
	if s < 0 {
		return -s
	}

	return s
}

func Pow[T constraints.Integer | constraints.Float](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Sqrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Sqrt(float64(s)))
}

func Cbrt[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Cbrt(float64(s)))
}

func Hypot[T constraints.Integer | constraints.Float](x, y T) T {
	return T(math.Hypot(float64(x), float64(y)))
}

func FilterNot[T any](ss []T, f func(T) bool) []T {
	var us []T
	for _, s := range ss {
		if !f(s) {
			us = append(us, s)
		}
	}

	return us
}

func Reduce[T any](ss []T, f func(T, T) T) T {
	if len(ss) == 0 {
		return *new(T)
	}

	result := ss[0]
	for _, s := range ss[1:] {
		result = f(result, s)
	}

	return result
}

func Drop[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[n:]
}

func Any[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if f(s) {
			return true
		}
	}

	return false
}

func All[T any](ss []T, f func(T) bool) bool {
	for _, s := range ss {
		if !f(s) {
			return false
		}
	}

	return true
}

func Shuffle[T any](ss []T) {
	for i := range ss {
		j := rand.Intn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}
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
	m := make(map[T]struct{})
	for _, s := range ss {
		m[s] = struct{}{}
	}

	for s := range m {
		ret = append(ret, s)
	}

	return
}

func UniqueUsing[T any](ss []T, f func(T) any) (ret []T) {
	m := make(map[any]T)
	for _, s := range ss {
		m[(f(s))] = s
	}

	for _, s := range m {
		ret = append(ret, s)
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
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return
}

func Map[T, U any](ss []T, f func(T) U) (ret []U) {
	for _, s := range ss {
		ret = append(ret, f(s))
	}

	return
}

func Contains[T constraints.Ordered](ss []T, s T) bool {
	for _, v := range ss {
		if v == s {
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

	for i := 0; i < n; i++ {
		ret = append(ret, ss[i])
	}

	return
}

func Average[T constraints.Integer | constraints.Float](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return Sum(ss) / T(len(ss))
}

func Reverse[T any](ss []T) (ret []T) {
	for i := len(ss) - 1; i >= 0; i-- {
		ret = append(ret, ss[i])
	}

	return
}

func Chunk[T any](ss []T, size int) (ret [][]T) {
	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}

		ret = append(ret, ss[i:end])
	}

	return
}
