package candy

import "golang.org/x/exp/constraints"

// Slice2Map converts a slice to a map where each element becomes a key with value true
func Slice2Map[M constraints.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}

// Slice2MapWithIndex converts a slice to a map where values are indices
func Slice2MapWithIndex[T comparable](list []T) map[T]int {
	if list == nil {
		return nil
	}

	result := make(map[T]int, len(list))
	for i, item := range list {
		result[item] = i
	}

	return result
}
