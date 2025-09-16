package candy

import "cmp"

// Slice2Map converts a slice to a map where each element becomes a key with value true
func Slice2Map[M cmp.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}

// SliceToMapGeneric converts a slice to a map where each element becomes a key with value true
func SliceToMapGeneric[T comparable](list []T) map[T]bool {
	if list == nil {
		return nil
	}

	result := make(map[T]bool, len(list))
	for _, item := range list {
		result[item] = true
	}

	return result
}

// SliceToMapWithValue converts a slice to a map with a custom value
func SliceToMapWithValue[T comparable, V any](list []T, value V) map[T]V {
	if list == nil {
		return nil
	}

	result := make(map[T]V, len(list))
	for _, item := range list {
		result[item] = value
	}

	return result
}

// SliceToMapWithIndex converts a slice to a map where values are indices
func SliceToMapWithIndex[T comparable](list []T) map[T]int {
	if list == nil {
		return nil
	}

	result := make(map[T]int, len(list))
	for i, item := range list {
		result[item] = i
	}

	return result
}