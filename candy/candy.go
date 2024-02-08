package candy

import (
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
)

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

func Sum[T constraints.Integer | constraints.Float](ss []T) (sum T) {
	for _, s := range ss {
		sum += s
	}

	return
}

func Average[T constraints.Integer | constraints.Float](ss []T) (avg T) {
	if len(ss) == 0 {
		return
	}

	sum := Sum(ss)
	avg = sum / T(len(ss))

	return
}

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

func Log[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Log(float64(s)))
}

func Log2[T constraints.Integer | constraints.Float](s T) T {
	return T(math.Log2(float64(s)))
}

func Each[T any](ss []T, f func(T)) {
	for _, s := range ss {
		f(s)
	}
}

func Map[T, U any](ss []T, f func(T) U) []U {
	us := make([]U, len(ss))
	for i, s := range ss {
		us[i] = f(s)
	}

	return us
}

func Filter[T any](ss []T, f func(T) bool) []T {
	var us []T
	for _, s := range ss {
		if f(s) {
			us = append(us, s)
		}
	}

	return us
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

func FirstOr[T any](ss []T, or T) T {
	if len(ss) > 0 {
		return ss[0]
	}

	return or
}

func First[T any](ss []T) (first T) {
	if len(ss) > 0 {
		first = ss[0]
	}

	return
}

func LastOr[T any](ss []T, or T) T {
	if len(ss) > 0 {
		return ss[len(ss)-1]
	}

	return or
}

func Last[T any](ss []T) (last T) {
	if len(ss) > 0 {
		last = ss[len(ss)-1]
	}

	return
}

func Top[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[:n]
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

func Reverse[T any](ss []T) []T {
	var us []T
	for i := len(ss) - 1; i >= 0; i-- {
		us = append(us, ss[i])
	}

	return us
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

func Random[T any](ss []T) T {
	if len(ss) == 0 {
		return *new(T)
	}

	return ss[rand.Intn(len(ss))]
}

func Shuffle[T any](ss []T) {
	for i := range ss {
		j := rand.Intn(i + 1)
		ss[i], ss[j] = ss[j], ss[i]
	}
}
